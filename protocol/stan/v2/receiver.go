/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package stan

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/Yangfisher1/ce-go-sdk/v2/binding"
	"github.com/nats-io/stan.go"
)

type msgErr struct {
	msg binding.Message
	err error
}

// Receiver implements protocol.Receiver for STAN subscriptions
type Receiver struct {
	incoming    chan msgErr
	messageOpts []MessageOption
}

func NewReceiver(opts ...ReceiverOption) (*Receiver, error) {
	r := &Receiver{
		incoming: make(chan msgErr),
	}

	if err := r.applyOptions(opts...); err != nil {
		return nil, err
	}

	return r, nil
}

// MsgHandler implements stan.MsgHandler
// This function is passed to the call to stan.Conn.Subscribe so that we can stream messages to be delivered
// via Receive()
func (r *Receiver) MsgHandler(msg *stan.Msg) {
	m, err := NewMessage(msg, r.messageOpts...)
	r.incoming <- msgErr{msg: m, err: err}
}

// Receive implements Receiver.Receive
// This should probably not be invoked directly by applications or library code, but instead invoked via
// Protocol.Receive
func (r *Receiver) Receive(ctx context.Context) (binding.Message, error) {
	select {
	case msgErr, ok := <-r.incoming:
		if !ok {
			return nil, io.EOF
		}
		return msgErr.msg, msgErr.err
	case <-ctx.Done():
		return nil, io.EOF
	}
}

func (r *Receiver) applyOptions(opts ...ReceiverOption) error {
	for _, fn := range opts {
		if err := fn(r); err != nil {
			return err
		}
	}
	return nil
}

// Consumer is responsible for managing STAN subscriptions and makes messages available via the Receiver interface.
//
// Consumer implements the following interfaces:
//
// - protocol.Opener
// - protocol.Closer
// - protocol.Receiver
type Consumer struct {
	Receiver

	Conn               stan.Conn
	Subject            string
	Subscriber         Subscriber
	UnsubscribeOnClose bool

	subscriptionOptions []stan.SubscriptionOption
	appliedSubOpts      *stan.SubscriptionOptions // only used locally, stored to avoid recomputing

	subMtx        sync.Mutex
	internalClose chan struct{}
	connOwned     bool // whether this consumer is responsible for closing the connection
}

func NewConsumer(clusterID, clientID, subject string, stanOpts []stan.Option, opts ...ConsumerOption) (*Consumer, error) {
	conn, err := stan.Connect(clusterID, clientID, stanOpts...)
	if err != nil {
		return nil, err
	}

	c, err := NewConsumerFromConn(conn, subject, opts...)
	if err != nil {
		if err2 := conn.Close(); err2 != nil {
			return nil, fmt.Errorf("failed to close conn: %s, when recovering from err: %w", err2, err)
		}
		return nil, err
	}

	return c, err
}

func NewConsumerFromConn(conn stan.Conn, subject string, opts ...ConsumerOption) (*Consumer, error) {
	c := &Consumer{
		Conn:          conn,
		Subject:       subject,
		Subscriber:    &RegularSubscriber{},
		internalClose: make(chan struct{}, 1),
	}

	err := c.applyOptions(opts...)
	if err != nil {
		return nil, err
	}

	receiverOps, err := c.createReceiverOptions()
	if err != nil {
		return nil, err
	}

	r, err := NewReceiver(receiverOps...)
	if err != nil {
		return nil, err
	}
	c.Receiver = *r

	return c, nil
}

// OpenInbound implements Opener.OpenInbound.
func (c *Consumer) OpenInbound(ctx context.Context) error {
	c.subMtx.Lock()
	defer c.subMtx.Unlock()

	// Subscribe
	sub, err := c.Subscriber.Subscribe(c.Conn, c.Subject, c.Receiver.MsgHandler, c.subscriptionOptions...)
	if err != nil {
		return err
	}

	// Wait until external or internal context done
	select {
	case <-ctx.Done():
	case <-c.internalClose:
	}

	if c.UnsubscribeOnClose {
		return sub.Unsubscribe()
	} else {
		return sub.Close()
	}
}

// Close implements Closer.Close.
// This method only closes the connection if the Consumer opened it. Subscriptions are closed/unsubscribed dependent
// on the UnsubscribeOnClose field.
func (c *Consumer) Close(_ context.Context) error {
	// Before closing, let's be sure OpenInbound completes
	// We send a signal to close and then we lock on subMtx in order
	// to wait OpenInbound to finish draining the queue
	c.internalClose <- struct{}{}
	c.subMtx.Lock()
	defer c.subMtx.Unlock()

	if c.connOwned {
		return c.Conn.Close()
	}

	close(c.internalClose)

	return nil
}

func (c *Consumer) applyOptions(opts ...ConsumerOption) error {
	for _, fn := range opts {
		if err := fn(c); err != nil {
			return err
		}
	}
	return nil
}

// createReceiverOptions builds an array of ReceiverOption used to configure the receiver.
func (c *Consumer) createReceiverOptions() ([]ReceiverOption, error) {
	// receivers need to know whether or not the subscription is configured in ManualAck-mode,
	// as such we must build the options in the same way as stan does since their API doesn't
	// expose this information
	subOpts, err := c.subOptionsLazy()
	if err != nil {
		return nil, err
	}

	opts := make([]ReceiverOption, 0)
	if subOpts.ManualAcks {
		opts = append(opts, WithMessageOptions(WithManualAcks()))
	}

	return opts, nil
}

// subOptionsLazy calculates the SubscriptionOptions based on an array of SubscriptionOption and stores the result on
// the struct to prevent repeated calculations
func (c *Consumer) subOptionsLazy() (*stan.SubscriptionOptions, error) {
	if c.appliedSubOpts != nil {
		return c.appliedSubOpts, nil
	}

	subOpts := stan.DefaultSubscriptionOptions
	for _, fn := range c.subscriptionOptions {
		err := fn(&subOpts)
		if err != nil {
			return nil, err
		}
	}

	c.appliedSubOpts = &subOpts

	return c.appliedSubOpts, nil
}

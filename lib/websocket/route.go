// Copyright 2018, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"bytes"
	"context"
	"fmt"
)

//
// RouteHandler is a function that will be called when registered method and
// target match with request.
//
type RouteHandler func(ctx context.Context, req *Request) (res Response)

type route struct {
	name    string
	childs  []*route
	handler RouteHandler
	isParam bool
}

//
// addChild to route.
//
func (r *route) addChild(isParam bool, name string) (c *route, err error) {
	c = r.getChild(isParam, name)
	if c != nil {
		return
	}
	if isParam {
		c = r.getChildAsParam()
		if c != nil {
			err = ErrRouteDupParam
			c = nil
			return
		}
	}
	c = &route{
		name:    name,
		isParam: isParam,
	}
	r.childs = append(r.childs, c)

	return
}

//
// getChild of current route which has the same isParam and name value.  It
// will return nil if not found.
//
func (r *route) getChild(isParam bool, name string) *route {
	for _, c := range r.childs {
		if isParam == c.isParam && name == c.name {
			return c
		}
	}
	return nil
}

//
// getChildAsParam return child route which type is parameter.
//
func (r *route) getChildAsParam() *route {
	for _, c := range r.childs {
		if c.isParam {
			return c
		}
	}
	return nil
}

//
// String return route representation as string.  This function is to prevent
// formatted print to print pointer to route as address.
//
func (r *route) String() (out string) {
	bb := _bbPool.Get().(*bytes.Buffer)
	bb.Reset()

	fmt.Fprintf(bb, "{name:%s", r.name)
	fmt.Fprintf(bb, " isParam:%v", r.isParam)
	fmt.Fprintf(bb, " handler:%v", r.handler)
	bb.WriteString(" childs:[")

	for x, c := range r.childs {
		if x > 0 {
			bb.WriteByte(' ')
		}
		fmt.Fprintf(bb, "%s", c.String())
	}

	bb.WriteString("]}")

	out = bb.String()

	_bbPool.Put(bb)

	return
}

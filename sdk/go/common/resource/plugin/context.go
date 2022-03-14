// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"context"
	"io/ioutil"

	"github.com/opentracing/opentracing-go"

	"github.com/mariospas/pulumi/sdk/v3/go/common/diag"
	"github.com/mariospas/pulumi/sdk/v3/go/common/diag/colors"
	"github.com/mariospas/pulumi/sdk/v3/go/common/util/rpcutil"
)

// Context is used to group related operations together so that
// associated OS resources can be cached, shared, and reclaimed as
// appropriate. It also carries shared plugin configuration.
type Context struct {
	Diag       diag.Sink // the diagnostics sink to use for messages.
	StatusDiag diag.Sink // the diagnostics sink to use for status messages.
	Host       Host      // the host that can be used to fetch providers.
	Pwd        string    // the working directory to spawn all plugins in.
	Root       string    // the root directory of the project.

	tracingSpan opentracing.Span // the OpenTracing span to parent requests within.
}

// NewContext allocates a new context with a given sink and host. Note
// that the host is "owned" by this context from here forwards, such
// that when the context's resources are reclaimed, so too are the
// host's.
func NewContext(d, statusD diag.Sink, host Host, cfg ConfigSource,
	pwd string, runtimeOptions map[string]interface{}, disableProviderPreview bool,
	parentSpan opentracing.Span) (*Context, error) {

	root := ""
	return NewContextWithRoot(d, statusD, host, cfg, pwd, root, runtimeOptions, disableProviderPreview, parentSpan)
}

// Variation of NewContext that also sets known project Root.
func NewContextWithRoot(d, statusD diag.Sink, host Host, cfg ConfigSource,
	pwd, root string, runtimeOptions map[string]interface{}, disableProviderPreview bool,
	parentSpan opentracing.Span) (*Context, error) {

	if d == nil {
		d = diag.DefaultSink(ioutil.Discard, ioutil.Discard, diag.FormatOptions{Color: colors.Never})
	}
	if statusD == nil {
		statusD = diag.DefaultSink(ioutil.Discard, ioutil.Discard, diag.FormatOptions{Color: colors.Never})
	}

	ctx := &Context{
		Diag:        d,
		StatusDiag:  statusD,
		Host:        host,
		Pwd:         pwd,
		tracingSpan: parentSpan,
	}
	if host == nil {
		h, err := NewDefaultHost(ctx, cfg, runtimeOptions, disableProviderPreview)
		if err != nil {
			return nil, err
		}
		ctx.Host = h
	}
	return ctx, nil
}

// Request allocates a request sub-context.
func (ctx *Context) Request() context.Context {
	// TODO[pulumi/pulumi#143]: support cancellation.
	return opentracing.ContextWithSpan(context.Background(), ctx.tracingSpan)
}

// Close reclaims all resources associated with this context.
func (ctx *Context) Close() error {
	if ctx.tracingSpan != nil {
		ctx.tracingSpan.Finish()
	}
	err := ctx.Host.Close()
	if err != nil && !rpcutil.IsBenignCloseErr(err) {
		return err
	}
	return nil
}

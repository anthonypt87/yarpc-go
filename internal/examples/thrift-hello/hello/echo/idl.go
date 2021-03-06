// Code generated by thriftrw v1.13.1. DO NOT EDIT.
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package echo

import "go.uber.org/thriftrw/thriftreflect"

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "echo",
	Package:  "go.uber.org/yarpc/internal/examples/thrift-hello/hello/echo",
	FilePath: "echo.thrift",
	SHA1:     "619e87984e33d89e2f63c235bf6b0c1003898504",
	Raw:      rawIDL,
}

const rawIDL = "service Hello {\n    EchoResponse echo(1:EchoRequest echo)\n}\n\nstruct EchoRequest {\n    1: required string message;\n    2: required i16 count;\n}\n\nstruct EchoResponse {\n    1: required string message;\n    2: required i16 count;\n}\n"

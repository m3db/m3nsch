// Copyright (c) 2016 Uber Technologies, Inc.
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

package xerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMultiErrorNoError(t *testing.T) {
	err := NewMultiError()
	require.Nil(t, err.FinalError())
	require.Equal(t, "", err.Error())
	require.True(t, err.Empty())
	require.Equal(t, 0, err.NumErrors())
}

func TestMultiErrorOneError(t *testing.T) {
	err := NewMultiError()
	err = err.Add(errors.New("foo"))
	final := err.FinalError()
	require.NotNil(t, final)
	require.Equal(t, "foo", final.Error())
	require.False(t, err.Empty())
	require.Equal(t, 1, err.NumErrors())
}

func TestMultiErrorMultipleErrors(t *testing.T) {
	err := NewMultiError()
	for _, errMsg := range []string{"foo", "bar", "baz"} {
		err = err.Add(errors.New(errMsg))
	}
	err = err.Add(nil)
	final := err.FinalError()
	require.NotNil(t, final)
	require.Equal(t, final.Error(), "foo\nbar\nbaz")
	require.False(t, err.Empty())
	require.Equal(t, 3, err.NumErrors())
}

func TestErrorsIsAnErrorAndFormatsErrors(t *testing.T) {
	errs := error(Errors{
		fmt.Errorf("some error: foo=2, bar=baz"),
		fmt.Errorf("some other error: foo=42, bar=qux"),
	})
	assert.Equal(t, "[<some error: foo=2, bar=baz>, "+
		"<some other error: foo=42, bar=qux>]", errs.Error())
}

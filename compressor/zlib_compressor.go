// Copyright 2022 <mzh.scnu@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package compressor

import (
	"bytes"
	"compress/zlib"
	"io"
)

// ZlibCompressor implements the Compressor interface
type ZlibCompressor struct {
}

// Zip .
func (ZlibCompressor) Zip(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	w := zlib.NewWriter(buf) // pointer
	defer w.Close()          // w is a pointer, so we can use w.Close() directly instead of using anonymous function to defer w.Close()
	_, err := w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Flush()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

// Unzip .
func (ZlibCompressor) Unzip(data []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	data, err = io.ReadAll(r) // replace ioutil.ReadAll(r) to io.ReadAll(r) in go1.16
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return nil, err
	}
	return data, nil
}

// Copyright 2022 <mzh.scnu@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package compressor

import (
	"bytes"
	"io"

	"github.com/golang/snappy"
)

// Snappy 是由 Google 开发的一种压缩库，专注于高压缩速度和合理的压缩比。Snappy 并不追求最小化的压缩率，而是优先保证较高的压缩和解压速度，使其非常适合需要快速处理大量数据的场景。
// SnappyCompressor implements the Compressor interface
type SnappyCompressor struct {
}

// Zip .
func (SnappyCompressor) Zip(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	w := snappy.NewBufferedWriter(buf)
	defer func() { // TODO: optimize. Whether it is necessary to use anonymous function to defer w.Close() because w is a pointer
		w.Close()
	}()
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
func (SnappyCompressor) Unzip(data []byte) ([]byte, error) {
	r := snappy.NewReader(bytes.NewBuffer(data))
	data, err := io.ReadAll(r) // replace ioutil.ReadAll(r) to io.ReadAll(r) in go1.16
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return nil, err
	}
	return data, nil
}

// Copyright 2022 <mzh.scnu@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package compressor

import (
	"bytes"
	"compress/gzip"
	"io"
)

// Zlib 是一个广泛使用的压缩库，基于 DEFLATE 算法，具有良好的压缩比和较快的速度。
// GzipCompressor implements the Compressor interface
type GzipCompressor struct {
}

// Zip. GzipCompressor object is not used in this func so omit the receiver name
func (GzipCompressor) Zip(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	w := gzip.NewWriter(buf)
	defer w.Close() // TODO: optimize. use anonymous function to defer w.Close() incase of modifing the w object
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

// Unzip.
func (GzipCompressor) Unzip(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewBuffer(data))
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

// Copyright 2022 <mzh.scnu@qq.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package codec

import (
	"encoding/binary"
	"io"
	"net"
)

// 若写入数据的长度为 0 ，此时sendFrame 函数会向IO流写入uvarint类型的 0 值；
// 若写入数据的长度大于 0 ，此时sendFrame 函数会向IO流写入uvarint类型的 len(data) 值，随后将该字节串的数据 data 写入IO流中。
func sendFrame(w io.Writer, data []byte) (err error) {
	var size [binary.MaxVarintLen64]byte

	if len(data) == 0 {
		n := binary.PutUvarint(size[:], uint64(0))
		if err = write(w, size[:n]); err != nil { // TODO: optimize
			return
		}
		return
	}

	n := binary.PutUvarint(size[:], uint64(len(data)))
	if err = write(w, size[:n]); err != nil {
		return
	}
	if err = write(w, data); err != nil {
		return
	}
	return
}

// 首先会向IO中读入uvarint类型的 size ，表示要接收数据的长度，随后将该从IO流中读取该 size 长度字节串.
// 由于 codec 层会传入一个bufio类型的结构体，bufio类型实现了有缓冲的IO操作，以便减少IO在用户态与内核态拷贝的次数
func recvFrame(r io.Reader) (data []byte, err error) {
	size, err := binary.ReadUvarint(r.(io.ByteReader))
	if err != nil {
		return nil, err
	}
	if size != 0 {
		data = make([]byte, size)
		if err = read(r, data); err != nil {
			return nil, err
		}
	}
	return data, nil
}

// write the body data to the writer.
// writer is an interface that wraps the basic Write method.
func write(w io.Writer, data []byte) error {
	for index := 0; index < len(data); {
		n, err := w.Write(data[index:])
		if _, ok := err.(net.Error); !ok {
			return err
		}
		index += n
	}
	return nil
}

// read the data from the reader and store it in the data slice.
func read(r io.Reader, data []byte) error {
	for index := 0; index < len(data); {
		n, err := r.Read(data[index:])
		if err != nil {
			if _, ok := err.(net.Error); !ok {
				return err
			}
		}
		index += n
	}
	return nil
}

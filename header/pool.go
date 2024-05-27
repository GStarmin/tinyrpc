// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package header

import "sync"

var (
	RequestPool  sync.Pool
	ResponsePool sync.Pool
)

func init() {
	// reduce the creating times of RequestHeader object
	RequestPool = sync.Pool{New: func() interface{} {
		return &RequestHeader{}
	}}
	// reduce the creating times of ReponseHeader object
	ResponsePool = sync.Pool{New: func() interface{} {
		return &ResponseHeader{}
	}}
}

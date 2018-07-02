// Copyright 2018 Hajime Hoshi
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

package main

import (
	"time"

	"github.com/gopherjs/gopherwasm/js"
)

func array() []float64 {
	return make([]float64, 16)
}

func heavyTask() float64 {
	arr := array()

	result := float64(0.0)

	for i := 0; i < 10000000; i++ {
		u0, v0, u1, v1 := float64(0.0), float64(0.0), float64(1.0), float64(1.0)
		arr[0] = u0
		arr[1] = v0
		arr[2] = u1
		arr[3] = u1

		arr[4] = u1
		arr[5] = v0
		arr[6] = u0
		arr[7] = u1

		arr[8] = u0
		arr[9] = v1
		arr[10] = u1
		arr[11] = u0

		arr[12] = u1
		arr[13] = v1
		arr[14] = u0
		arr[15] = u0

		result += arr[i % 16]
	}
	return result
}

var now func() int64

func init() {
	if js.Global() != js.Null() {
		now = func() int64 {
			arr := js.Global().Get("process").Call("hrtime")
			return int64(arr.Index(0).Int() * 1e9 + arr.Index(1).Int())
		}
		return
	}
	now = func() int64 {
		return int64(time.Now().UnixNano())
	}
}

func main() {
	t1 := now()
	x := heavyTask()
	t2 := now()
	println("result:", x)
	println("time [ns]:", int32(t2-t1))
}

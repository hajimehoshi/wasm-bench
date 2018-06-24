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

function array() {
  return new Float32Array(16);
}

function heavyTask() {
  const arr = array();
  result = 0.0;
  for (let i = 0; i < 10000000; i++) {
    const u0 = 0.0;
    const v0 = 0.0;
    const u1 = 1.0;
    const v1 = 1.0;

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
  return result;
}

function main() {
  const t1 = process.hrtime();
  const x = heavyTask();
  const diff = process.hrtime(t1);
  console.log('result:', x);
  console.log('time [ns]:', diff[0] * 1e9 + diff[1]);
}

main();

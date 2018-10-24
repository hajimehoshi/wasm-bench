echo 'Go Native'
go run main.go
echo

echo 'C++ Native (-Os)'
g++ -std=c++14 -Os -o out.c++ main.cc
./out.c++
echo

echo 'Wasm (Go)'
GOOS=js GOARCH=wasm go build -o out.go.wasm .
./go_js_wasm_exec out.go.wasm
echo

echo 'Wasm (Go) (-Os by wasm-opt)'
wasm-opt -Os out.go.wasm -o out.go.os.wasm
./go_js_wasm_exec out.go.os.wasm
echo

echo 'Wasm (Go) (-O4 by wasm-opt)'
wasm-opt -O4 out.go.wasm -o out.go.o4.wasm
./go_js_wasm_exec out.go.o4.wasm
echo

echo 'GopherJS'
gopherjs build -o out.gopherjs.js .
node out.gopherjs.js
echo

echo 'JavaScript'
node main.js
echo

echo 'Wasm (C++) (-O3)'
emcc main.cc -std=c++14 -O3 -o out.emcc.js -g -s WASM=1
node out.emcc.js
echo

echo 'Go Native'
go run main.go
echo

echo 'C++ Native (-Os)'
g++ -std=c++14 -Os -o out.c++ main.cc
./out.c++
echo

echo 'Wasm (Go)'
GOOS=js GOARCH=wasm ~/go-code/bin/go build -o out.go.wasm .
./go_js_wasm_exec out.go.wasm
echo

echo 'GopherJS'
gopherjs build -o out.gopherjs.js .
node out.gopherjs.js
echo

echo 'JavaScript'
node main.js
echo

echo 'Wasm (C++) (-Os)'
emcc main.cc -std=c++14 -Os -s -o out.emcc.js WASM=1
node out.emcc.js
echo

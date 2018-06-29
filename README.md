# Result on MacBook Pro 2014 Mid (2018-06-29)

```
$ ./run.sh
Go Native
result: +5.000000e+006
time [ns]: 46962513

C++ Native (-Os)
result: 5000000.000000
time [ns]: 11941360

Wasm (Go)
result: +5.000000e+006
time [ns]: 226033921

Wasm (Go) (-Os by wasm-opt)
result: +5.000000e+006
time [ns]: 176618698

Wasm (Go) (-O4 by wasm-opt)
result: +5.000000e+006
time [ns]: 155280898

GopherJS
result: 5000000
time [ns]: 143573587

JavaScript
result: 5000000
time [ns]: 74950863

Wasm (C++) (-Os)
result: 5000000.000000
time [ns]: 62829009
```

# Result on MacBook Pro 2014 Mid (2018-06-25)

```
$ ./run.sh
Go Native
result: +5.000000e+006
time [ns]: 50207750

C++ Native (-Os)
result: 5000000.000000
time [ns]: 12605950

Wasm (Go)
result: +5.000000e+006
time [ns]: 257062926

GopherJS
result: 5000000
time [ns]: 154164104

JavaScript
result: 5000000
time [ns]: 95847509

Wasm (C++) (-Os)
result: 5000000.000000
time [ns]: 70553509
```

# Result on MacBook Pro 2014 Mid (2018-06-24)

```
$ ./run.sh
Go Native
result: +5.000000e+006
time [ns]: 52339041

C++ Native (-Os)
result: 5000000.000000
time [ns]: 14312885

Wasm (Go)
result: +5.000000e+006
time [ns]: 243000000

GopherJS
result: 5000000
time [ns]: 158000000

JavaScript
result: 5000000
time [ns]: 89701023

Wasm (C++) (-Os)
result: 5000000.000000
time [ns]: 78170506
```

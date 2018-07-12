# Result on MacBook Pro 2014 Mid (2018-07-13)

## Versions

### Go

```
go version devel +9fa988547a Thu Jul 5 07:21:50 2018 +0000 darwin/amd64
```

### Node

```
v10.6.0
```

### Emscripten

```
1.38.6
```

## f32

```
Go Native
result: +5.000000e+006
time [ns]: 46286663

C++ Native (-Os)
result: 5000000.000000
time [ns]: 11765504

Wasm (Go)
result: +5.000000e+006
time [ns]: 182948167

Wasm (Go) (-Os by wasm-opt)
result: +5.000000e+006
time [ns]: 148658133

Wasm (Go) (-O4 by wasm-opt)
result: +5.000000e+006
time [ns]: 128597310

GopherJS
result: 5000000
time [ns]: 129638379

JavaScript
result: 5000000
time [ns]: 61945516

Wasm (C++) (-Os)
result: 5000000.000000
time [ns]: 66390088
```

## f64

```
Go Native
result: +5.000000e+006
time [ns]: 47015580

C++ Native (-Os)
result: 5000000.000000
time [ns]: 41005199

Wasm (Go)
result: +5.000000e+006
time [ns]: 161097461

Wasm (Go) (-Os by wasm-opt)
result: +5.000000e+006
time [ns]: 153063944

Wasm (Go) (-O4 by wasm-opt)
result: +5.000000e+006
time [ns]: 128250647

GopherJS
result: 5000000
time [ns]: 112628008

JavaScript
result: 5000000
time [ns]: 61670484

Wasm (C++) (-Os)
result: 5000000.000000
time [ns]: 62569073
```

# Result on MacBook Pro 2014 Mid (2018-07-02)

## f32

```
Go Native
result: +5.000000e+006
time [ns]: 46503928

C++ Native (-Os)
result: 5000000.000000
time [ns]: 12024699

Wasm (Go)
result: +5.000000e+006
time [ns]: 223048008

Wasm (Go) (-Os by wasm-opt)
result: +5.000000e+006
time [ns]: 179819890

Wasm (Go) (-O4 by wasm-opt)
result: +5.000000e+006
time [ns]: 159952173

GopherJS
result: 5000000
time [ns]: 149482927

JavaScript
result: 5000000
time [ns]: 73622416

Wasm (C++) (-Os)
result: 5000000.000000
time [ns]: 67350223
```

## f64

```
Go Native
result: +5.000000e+006
time [ns]: 46859944

C++ Native (-Os)
result: 5000000.000000
time [ns]: 40685821

Wasm (Go)
result: +5.000000e+006
time [ns]: 190058473

Wasm (Go) (-Os by wasm-opt)
result: +5.000000e+006
time [ns]: 173105512

Wasm (Go) (-O4 by wasm-opt)
result: +5.000000e+006
time [ns]: 154126142

GopherJS
result: 5000000
time [ns]: 142118787

JavaScript
result: 5000000
time [ns]: 73518466

Wasm (C++) (-Os)
result: 5000000.000000
time [ns]: 66163768
```

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

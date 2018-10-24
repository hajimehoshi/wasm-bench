# Result on MacBook Pro 2014 Mid (2018-10-25)

## Versions

### Go

```
go version go1.11.1 darwin/amd64
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
time [ns]: 47538000

C++ Native (-Os)
result: 5000000.000000
time [ns]: 13612005

Wasm (Go)
result: +5.000000e+006
time [ns]: 187528488

Wasm (Go) (-Os by wasm-opt)
result: +5.000000e+006
time [ns]: 147503206

Wasm (Go) (-O4 by wasm-opt)
result: +5.000000e+006
time [ns]: 133056493

GopherJS
result: 5000000
time [ns]: 133925725

JavaScript
result: 5000000
time [ns]: 65524751

Wasm (C++) (-O3)
result: 5000000.000000
time [ns]: 66015871
```

## f64

```
Go Native
result: +5.000000e+006
time [ns]: 47116000

C++ Native (-Os)
result: 5000000.000000
time [ns]: 41781047

Wasm (Go)
result: +5.000000e+006
time [ns]: 169768051

Wasm (Go) (-Os by wasm-opt)
result: +5.000000e+006
time [ns]: 145240577

Wasm (Go) (-O4 by wasm-opt)
result: +5.000000e+006
time [ns]: 129833968

GopherJS
result: 5000000
time [ns]: 117594827

JavaScript
result: 5000000
time [ns]: 63437460

Wasm (C++) (-O3)
result: 5000000.000000
time [ns]: 64420889
```

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
time [ns]: 46852896

C++ Native (-Os)
result: 5000000.000000
time [ns]: 11827945

Wasm (Go)
result: +5.000000e+006
time [ns]: 192059527

Wasm (Go) (-Os by wasm-opt)
result: +5.000000e+006
time [ns]: 141913384

Wasm (Go) (-O4 by wasm-opt)
result: +5.000000e+006
time [ns]: 142049012

GopherJS
result: 5000000
time [ns]: 129323688

JavaScript
result: 5000000
time [ns]: 74149606

Wasm (C++) (-O3)
result: 5000000.000000
time [ns]: 64772805
```

## f64

```
Go Native
result: +5.000000e+006
time [ns]: 46454333

C++ Native (-Os)
result: 5000000.000000
time [ns]: 40871053

Wasm (Go)
result: +5.000000e+006
time [ns]: 160427612

Wasm (Go) (-Os by wasm-opt)
result: +5.000000e+006
time [ns]: 142731077

Wasm (Go) (-O4 by wasm-opt)
result: +5.000000e+006
time [ns]: 130341177

GopherJS
result: 5000000
time [ns]: 111787121

JavaScript
result: 5000000
time [ns]: 63478679

Wasm (C++) (-O3)
result: 5000000.000000
time [ns]: 64402912
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

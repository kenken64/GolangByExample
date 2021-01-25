## Chapter 1

### Make map

https://blog.golang.org/maps

```
m := make(map[string]int)
```

### Short hand declaration

```
m := 0.0
```

### Underscore (Blank identifier)

```
_, err := test()

```

### Channel operator <- arrow

```
ch chan<-string
```

```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

## Chapter 2

Keyword 

```
break       default     func    interface   select
case        defer       go      map         struct
chan        else        goto    package     switch
const       fallthorugh if      range       type
continue    for         import  return      var
```

Constants 
```
true false iota nil
```

Types 
```
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex128 complex64
bool byte rune string error
```

Functions
```
make len cap new append copy close delete
complex real imag
panic recover
```

### Pointers

```
x:=1
p:=&x
fmt.Printf(*p)
*p= 2
fmt.Printf(x)
```

### fmt printF arg

fmt (Args) | Describe
------------ | -------------
%d | decimal integer 
%x, | %o, %b integer in hexadecimal, octal, binary 
%f, | %g, %e floating-point number: 3.141593 3.141592653589793 3.141593e+00 
%t | boolean: true or false 
%c | rune (Unicode code point) 
%s | string 
%q | quoted string "abc" or rune 'c' 
%v | any value in a natural format 
%T | type of any value 
%% | literal percent sign (no operand)


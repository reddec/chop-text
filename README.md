# Chop string

Very simple library to chop string to slice


Examples:


```go
ch:=Chopper{Slices: 3, Sep: '/', SliceSize: 3}
println(ch.Chop("abcdefxyz123"))
// will print: abc/def/xyz123
```

[Documentation](http://godoc.org/github.com/reddec/chop-text)
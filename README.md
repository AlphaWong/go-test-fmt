# Objective
Test hex conversion benchmark in <code>fmt.Sprintf("%x",secret)</code> and <code>hex.EncodeToString(secret)</code>

# Expected
Same benchmark.

# Actual
<code>hex.EncodeToString(secret)</code> always faster then fmt.

# Run 
```bash
go test -run Benchmark -test.bench ConvertHex -test.benchmem
```
# Result
```bash
BenchmarkFmtConvertHex-4         	 2000000	       720 ns/op	     272 B/op	       5 allocs/op
BenchmarkEncodeHexConvertHex-4   	 3000000	       580 ns/op	     304 B/op	       5 allocs/op
PASS
ok  	go-test-fmt	4.533s
```

# Go version
go 1.8.3

# Code
```go
func BenchmarkFmtConvertHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmtConvertHex(getDummySecretMessage())
	}
}

func BenchmarkEncodeHexConvertHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeHexConvertHex(getDummySecretMessage())
	}
}

func fmtConvertHex(secret []byte) string{
  // https://golang.org/pkg/crypto/sha256/#New
	return fmt.Sprintf("%x",secret)
}

func encodeHexConvertHex(secret []byte) string{
	return hex.EncodeToString(secret)
}

func getDummySecretMessage() []byte{
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	return h.Sum(nil)
}
```
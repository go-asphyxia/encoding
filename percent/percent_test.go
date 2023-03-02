package percent

import (
	"log"
	"net/url"
	"testing"

	"github.com/go-asphyxia/core/bytes"
)

var (
	simpleE       = "%D0%BF%D1%80%D0%BE%D1%81%D1%82%D0%B0%D1%8F%20%D1%81%D1%82%D1%80%D0%BE%D0%BA%D0%B0"
	simpleEBuffer = bytes.Buffer(simpleE)
	simpleD       = "простая строка"
	simpleDBuffer = bytes.Buffer(simpleD)
)

func TestPercent(t *testing.T) {
	decoded := Decode(simpleEBuffer)

	log.Println(simpleE)
	log.Println(Encode(decoded).String())

	decodedStd, _ := url.QueryUnescape(simpleE)

	log.Println(simpleE)
	log.Println(url.QueryEscape(decodedStd))
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Decode(simpleEBuffer)
	}
}

func BenchmarkDecodeStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = url.QueryUnescape(simpleE)
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Encode(simpleDBuffer)
	}
}

func BenchmarkEncodeStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = url.QueryEscape(simpleD)
	}
}

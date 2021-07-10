package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAji(t *testing.T) {
	result := HelloWorld("Aji")
	if result != "Hello Aji" {
		//ksh error tapi tetep lanjut execute code selanjut nya
		// t.Fail()

		//sama kaya Fail() cuma bs kasih argument
		t.Error("Result nya harus nya Hai Aji")
	}

	fmt.Println("Test Hello World Aji Done")
}

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")
	if result != "Hello Budi" {
		// ksh error dan ga lanjut execute code selanjut nya
		// t.FailNow()

		// sama kayak FailNow() cuma bs dikasih argumen
		t.Fatal("Result nya harus nya Hai Budi")
	}
	fmt.Println("Test Hello World Budi Done")
}

//PAKE LIBRARY TESTIFY BUAT UNIT TEST
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Aji")
	assert.Equal(t, "Hello Aji", result, "Result nya harus nya Hai Aji")
	fmt.Println("Test Hello World Budi With Assert Equal Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Aji")
	require.Equal(t, "Hello Aji", result, "Result nya harus nya Hai Aji")
	fmt.Println("Test Hello World Budi With Require Equal Done")
}

func TestSkip(t *testing.T) {
	//Buat kondisi spesifik jika pengen nge skip code unit tes bs pake skip()
	if runtime.GOOS == "windows" {
		t.Skip("Unit Test Gabisa Jalan di MAC OS")
	}
	result := HelloWorld("Aji")
	require.Equal(t, "Hello Aji", result, "Result nya harus nya Hai Aji")
	fmt.Println("Test Hello World Budi With Require Equal Done")
}

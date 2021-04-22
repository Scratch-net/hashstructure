package hashstructure

import (
	"fmt"
)

func ExampleHash() {
	type ComplexStruct struct {
		Name     string
		Age      uint
		Metadata map[string]interface{}
	}

	v := ComplexStruct{
		Name: "mitchellh",
		Age:  64,
		Metadata: map[string]interface{}{
			"car":      true,
			"location": "California",
			"siblings": []string{"Bob", "John"},
		},
	}

	hash, err := Hash(v, FormatV2, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d", hash)
	// Output:
	// [59 211 209 211 148 28 150 169 196 203 223 172 124 19 44 68 61 179 84 224 195 75 194 37 64 184 169 150 146 17 214 37]
}

func ExampleHash_v1() {
	type ComplexStruct struct {
		Name     string
		Age      uint
		Metadata map[string]interface{}
	}

	v := ComplexStruct{
		Name: "mitchellh",
		Age:  64,
		Metadata: map[string]interface{}{
			"car":      true,
			"location": "California",
			"siblings": []string{"Bob", "John"},
		},
	}

	hash, err := Hash(v, FormatV1, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d", hash)
	// Output:
	// [98 170 206 140 146 246 234 156 70 153 166 45 160 60 8 193 69 142 187 240 51 18 16 168 160 230 180 53 54 197 156 104]
}

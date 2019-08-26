package populating

import (
	"github.com/kateGlebova/seaports-catalogue/client-api/managing"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonService_Populate(t *testing.T) {
	populator := jsonService{fileName: "test.json", manager: managing.MockService{}}
	file, err := os.Open(populator.fileName)
	if err != nil {
		t.Fatal(err)
	}
	populator.file = file
	defer populator.file.Close()

	err = populator.Populate()
	assert.NoError(t, err)
}

func BenchmarkJsonService_Populate(b *testing.B) {
	populator := jsonService{fileName: "ports.json", manager: managing.MockService{}}
	file, err := os.Open(populator.fileName)
	if err != nil {
		b.Fatal(err)
	}
	populator.file = file
	defer populator.file.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		populator.Populate()
	}
}

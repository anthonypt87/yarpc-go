package buffer

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuffers(t *testing.T) {
	var wg sync.WaitGroup
	for g := 0; g < 10; g++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				buf := Get()
				assert.Zero(t, buf.Len(), "Expected truncated buffer")

				bytesOfNoise := make([]byte, rand.Intn(5000))
				_, err := rand.Read(bytesOfNoise)
				assert.NoError(t, err, "Unexpected error from rand.Read")
				_, err = buf.Write(bytesOfNoise)
				assert.NoError(t, err, "Unexpected error from buffer.Write")

				assert.Equal(t, buf.Len(), len(bytesOfNoise), "Expected same buffer size")

				Put(buf)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

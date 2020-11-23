package legacybech32

import (
	"math/rand"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/stretchr/testify/require"
)

func BenchmarkBech32ifyPubKey(b *testing.B) {
	pkBz := make([]byte, ed25519.PubKeySize)
	pk := &ed25519.PubKey{Key: pkBz}
	rng := rand.New(rand.NewSource(time.Now().Unix()))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rng.Read(pk.Key)
		b.StartTimer()

		_, err := Bech32ifyPubKey(Bech32PubKeyTypeConsPub, pk)
		require.NoError(b, err)
	}
}

func BenchmarkGetPubKeyFromBech32(b *testing.B) {
	pkBz := make([]byte, ed25519.PubKeySize)
	pk := &ed25519.PubKey{Key: pkBz}
	rng := rand.New(rand.NewSource(time.Now().Unix()))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rng.Read(pk.Key)

		pkStr, err := Bech32ifyPubKey(Bech32PubKeyTypeConsPub, pk)
		require.NoError(b, err)

		b.StartTimer()
		pk2, err := GetPubKeyFromBech32(Bech32PubKeyTypeConsPub, pkStr)
		require.NoError(b, err)
		require.Equal(b, pk, pk2)
	}
}
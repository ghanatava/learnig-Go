package palindrome

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandomPalindrome(t *testing.T){
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed :%d",seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i<1000;i++{
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

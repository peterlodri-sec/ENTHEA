package languages

import "testing"

func TestAllLanesRoundTrip(t *testing.T) {
	for _, l := range All() {
		ok, err := RoundTrip(l)
		if err != nil {
			t.Fatalf("%s: %v", l.Name, err)
		}
		if !ok {
			t.Errorf("%s: wire round-trip mismatch", l.Name)
		}
	}
}

func TestBridgePairDistinct(t *testing.T) {
	es, hu := BridgePair()
	if es == "" || hu == "" {
		t.Fatal("bridge lanes must be present")
	}
	if es == hu {
		t.Error("esperanto and hungarian frames must differ")
	}
}

func TestByName(t *testing.T) {
	if _, err := ByName("esperanto"); err != nil {
		t.Fatalf("esperanto lane missing: %v", err)
	}
	if _, err := ByName("magyar"); err != nil {
		t.Fatalf("hungarian native name missing: %v", err)
	}
	if _, err := ByName("klingon"); err == nil {
		t.Fatal("unknown lane must error")
	}
}

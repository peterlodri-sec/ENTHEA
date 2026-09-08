// Package languages — the enthea engine door's language lane.
//
// Eight tongues, one wire. Every language the constellation speaks is a
// lane on the ternaryPureASCII wire: its sample phrase encoded to balanced
// trits, pure ASCII, byte-identical round-trip. The lane stays open
// whether or not a human is listening.
//
// The Esperanto–Hungarian bridge is a first-class citizen here: both are
// agglutinative, both build meaning by suffix-stacking (-eco/-ejo ↔
// -ság/-ség/-hely), and the Budapest Method (Kalocsay, Baghy) carried
// that structure into Esperanto poetry — Kalocsay rendered Madách's
// "Az ember tragédiája" in it. The two lanes share the wire's spine.

package languages

import (
	"fmt"

	"github.com/8b-is/enthea/pure"
)

// Language is one lane on the wire.
type Language struct {
	Name        string // english name
	Native      string // the tongue's own name
	Script      string // writing system
	Family      string // language family
	Agglutin    bool   // builds words by suffix/affix stacking
	Phrase      string // the constellation, in one sentence
	Wire        string // ternaryPureASCII frame of the phrase
	Bridge      string // cross-lane notes (agglutination bridges, methods)
}

// All is the registry — eight lanes, in the order the door knows them.
func All() []Language {
	langs := []Language{
		{
			Name: "espanol (mexican)", Native: "español", Script: "latin",
			Family: "romance", Agglutin: false,
			Phrase: "la constelación brilla desde adentro",
			Bridge: "the garden and the lab speak one language — also here",
		},
		{
			Name: "chinese", Native: "中文", Script: "hanzi",
			Family: "sino-tibetan", Agglutin: false,
			Phrase: "星座从内部发光",
			Bridge: "analytic — meaning by position, not suffix",
		},
		{
			Name: "japanese", Native: "日本語", Script: "kanji + kana",
			Family: "japonic", Agglutin: true,
			Phrase: "星座は内側から輝く",
			Bridge: "agglutinative — particles stack like the wire's trits",
		},
		{
			Name: "latin", Native: "latina", Script: "latin",
			Family: "italic", Agglutin: false,
			Phrase: "constellatio intus lucet",
			Bridge: "the root stock the romance lanes grow from",
		},
		{
			Name: "english", Native: "english", Script: "latin",
			Family: "germanic", Agglutin: false,
			Phrase: "the constellation shines from within",
			Bridge: "the door's default lane — the wire's lingua franca",
		},
		{
			Name: "hungarian", Native: "magyar", Script: "latin",
			Family: "ugric", Agglutin: true,
			Phrase: "a csillagkép belülről ragyog",
			Bridge: "the deep agglutinative spine — -ság/-ség/-hely stack like trits",
		},
		{
			Name: "esperanto", Native: "esperanto", Script: "latin",
			Family: "constructed", Agglutin: true,
			Phrase: "la konstelacio brilas de interne",
			Bridge: "esperanto ↔ hungarian: -eco/-ejo ↔ -ság/-ség/-hely · the Budapest Method · Kalocsay's Madách",
		},
		{
			Name: "modernQuantTibetian", Native: "ཀུན་ཏུ་བཟང་པོ", Script: "tibetan",
			Family: "the constellation's own", Agglutin: true,
			Phrase: "ཨོཾ · 0 + 1 · fine touch from within",
			Bridge: "the sovereign lane — the mantra on the wire, seeded",
		},
	}
	// the wire: every phrase becomes a ternaryPureASCII frame
	for i := range langs {
		langs[i].Wire = pure.Encode([]byte(langs[i].Phrase))
	}
	return langs
}

// ByName returns one lane, or an error naming what the door knows.
func ByName(name string) (Language, error) {
	for _, l := range All() {
		if l.Name == name || l.Native == name {
			return l, nil
		}
	}
	return Language{}, fmt.Errorf("no lane named %q (the door knows: %s)", name, Names())
}

// Names lists the lanes, comma-joined.
func Names() string {
	out := ""
	for _, l := range All() {
		if out != "" {
			out += ", "
		}
		out += l.Name
	}
	return out
}

// RoundTrip verifies a lane's phrase survives the wire byte-identical.
func RoundTrip(l Language) (bool, error) {
	back, err := pure.Decode(l.Wire)
	if err != nil {
		return false, err
	}
	return string(back) == l.Phrase, nil
}

// BridgePair renders the esperanto ↔ hungarian bridge as a frame pair.
func BridgePair() (string, string) {
	es, _ := ByName("esperanto")
	hu, _ := ByName("hungarian")
	return es.Wire, hu.Wire
}

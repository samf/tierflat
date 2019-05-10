package tierflat

import (
	"fmt"
	"sort"
)

// Tier is a single tier to be used in a Tiered type.
type Tier struct {
	Name   string
	Units  int
	Abbrev string
	Amount int
}

// Tiered is a type for representing a system of tiers. It should be created via the
// MakeTiered call.
type Tiered struct {
	tiers  []Tier
	sorted bool
}

// MakeTiered creates a well-ordered Tiered out of a list of tiers.
func MakeTiered(tiers ...Tier) Tiered {
	t := Tiered{
		tiers: append([]Tier{}, tiers...),
	}

	t.sort()

	return t
}

// From takes a tiered as a prototype, and converts a flat amount
// into a new Tiered.
func (t Tiered) From(amount int) Tiered {
	if !t.sorted {
		t.sort()
	}

	tiers := t.tiers
	t.tiers = make([]Tier, len(tiers))
	for i := range tiers {
		t.tiers[i] = Tier{
			Name:   tiers[i].Name,
			Units:  tiers[i].Units,
			Abbrev: tiers[i].Abbrev,
			Amount: amount / tiers[i].Units,
		}

		amount %= tiers[i].Units
	}

	return t
}

func (t Tiered) String() string {
	var out string

	if !t.sorted {
		t.sort()
	}

	for _, ti := range t.tiers {
		if ti.Amount != 0 {
			out += fmt.Sprintf("%v%v", ti.Amount, ti.Abbrev)
		}
	}
	if out == "" {
		last := len(t.tiers) - 1
		out = fmt.Sprintf("0%v", t.tiers[last].Abbrev)
	}

	return out
}

func (t *Tiered) sort() {
	sort.Slice(t.tiers, func(i, j int) bool {
		return t.tiers[j].Units < t.tiers[i].Units
	})

	t.sorted = true
}

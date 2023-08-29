package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForMaldives(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m15.039 6.4H16.96C21.309 5.62 26.467 4 32 4s10.691 1.62 15.039 4.4m-4.223 40.947a18.62 18.62 0 0 1-3.605.354c-9.957 0-18.027-7.925-18.027-17.7c0-9.773 8.07-17.7 18.027-17.7c1.231 0 2.441.123 3.605.355c-8.23 1.64-14.421 8.782-14.421 17.344s6.191 15.704 14.421 17.347M16.96 55.6h30.079C42.691 58.38 37.533 60 32 60s-10.691-1.62-15.04-4.4"/>`),
		g.Group(children),
	)
}
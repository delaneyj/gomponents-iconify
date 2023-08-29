package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCzechia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm0 2c15.439 0 28 12.561 28 28H33.414l-20.48 20.48a29.01 29.01 0 0 1-1.414-1.414L30.586 32L11.52 12.934C16.635 7.443 23.922 4 32 4z"/>`),
		g.Group(children),
	)
}
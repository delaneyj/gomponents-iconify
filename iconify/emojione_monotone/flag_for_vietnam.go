package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForVietnam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm9.724 43.767L32 38.847l-9.726 6.92l3.674-11.255l-9.682-7.183h12.102L32 16.267l3.739 11.063h11.994l-9.682 7.183l3.673 11.254z"/>`),
		g.Group(children),
	)
}
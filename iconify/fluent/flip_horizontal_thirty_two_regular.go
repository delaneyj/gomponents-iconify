package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipHorizontalThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M29.842 27.54A1 1 0 0 1 29 28H18a1 1 0 0 1-1-1V3a1 1 0 0 1 1.91-.417l11 24a1 1 0 0 1-.068.957ZM19 7.582V26h8.442L19 7.582ZM3 28a1 1 0 0 1-.91-1.417l11-24A1 1 0 0 1 15 3v24a1 1 0 0 1-1 1H3Z"/>`),
		g.Group(children),
	)
}
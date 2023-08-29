package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRotaryStraight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 13a3 3 0 1 0 6 0a3 3 0 1 0-6 0m3 3v5m0-18v7M9 7l4-4l4 4"/>`),
		g.Group(children),
	)
}
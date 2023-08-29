package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YCombinator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm4.5 4l3.5 6v5h2v-5l3.5-6h-2L16 15.281L13.5 11z"/>`),
		g.Group(children),
	)
}
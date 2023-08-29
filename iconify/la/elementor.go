package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elementor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm4 4v10h2V11h-2zm4 0v2h6v-2h-6zm0 4v2h6v-2h-6zm0 4v2h6v-2h-6z"/>`),
		g.Group(children),
	)
}
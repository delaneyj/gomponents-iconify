package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 7v2h26V7zm0 4v2h26v-2zm0 4v2h26v-2zm0 4v2h26v-2zm0 4v2h26v-2z"/>`),
		g.Group(children),
	)
}
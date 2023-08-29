package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19v-9l4-4h1V4h2v2h9l4 4v9H2Zm14-2h4v-6.175l-2-2l-2 2V17ZM4 17h10v-5H4v5Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewComfyAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 10V2h8v8H2Zm0 12v-8h8v8H2Zm12-12V2h8v8h-8Zm0 12v-8h8v8h-8Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18V6h20v12h-2v-2h-4v2h-2V8H4v10H2Zm14-8h4V8h-4v2Zm0 4h4v-2h-4v2Z"/>`),
		g.Group(children),
	)
}
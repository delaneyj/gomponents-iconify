package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HPlusMobiledata(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17V7h2v4h6V7h2v10h-2v-4H6v4H4Zm14-2v-2h-2v-2h2V9h2v2h2v2h-2v2h-2Z"/>`),
		g.Group(children),
	)
}
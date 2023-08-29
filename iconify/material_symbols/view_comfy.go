package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewComfy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 11V4h20v7H2Zm8 9v-7h12v7H10Zm-8 0v-7h6v7H2Z"/>`),
		g.Group(children),
	)
}
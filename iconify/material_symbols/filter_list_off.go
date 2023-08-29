package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterListOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.775 22.625l-18.4-18.4L2.8 2.8l18.4 18.4l-1.425 1.425ZM15.825 13l-2-2H18v2h-2.175Zm-5-5l-2-2H21v2H10.825ZM10 18v-2h4v2h-4Zm-4-5v-2h4.15v2H6ZM3 8V6h2.15v2H3Z"/>`),
		g.Group(children),
	)
}
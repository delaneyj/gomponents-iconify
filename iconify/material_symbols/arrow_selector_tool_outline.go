package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSelectorToolOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 13.75L9.975 11h4.25L8 6.1v7.65ZM13.775 22l-3.625-7.8L6 20V2l14 11h-7.1l3.6 7.725L13.775 22Zm-3.8-11Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelRightGallerySixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 5.5A2.5 2.5 0 0 0 11.5 3h-7A2.5 2.5 0 0 0 2 5.5v5A2.5 2.5 0 0 0 4.5 13h7a2.5 2.5 0 0 0 2.5-2.5v-5Zm-1 0V6h-3V4h1.5A1.5 1.5 0 0 1 13 5.5ZM10 9V7h3v2h-3Zm0 1h3v.5a1.5 1.5 0 0 1-1.5 1.5H10v-2ZM9 4v8H4.5A1.5 1.5 0 0 1 3 10.5v-5A1.5 1.5 0 0 1 4.5 4H9Z"/>`),
		g.Group(children),
	)
}
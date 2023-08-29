package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelTopGallerySixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 3A2.5 2.5 0 0 0 2 5.5v5A2.5 2.5 0 0 0 4.5 13h7a2.5 2.5 0 0 0 2.5-2.5v-5A2.5 2.5 0 0 0 11.5 3h-7ZM3 5.5A1.5 1.5 0 0 1 4.5 4H6v3H3V5.5ZM3 8h10v2.5a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 10.5V8Zm10-1h-3V4h1.5A1.5 1.5 0 0 1 13 5.5V7ZM9 4v3H7V4h2Z"/>`),
		g.Group(children),
	)
}
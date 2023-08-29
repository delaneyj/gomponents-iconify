package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashOnOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 15.6l3.2-4.6h-2.85l2-7H9v8h3v3.6ZM10 22v-8H7V2h10l-2 7h4l-9 13Zm2-10H9h3Z"/>`),
		g.Group(children),
	)
}
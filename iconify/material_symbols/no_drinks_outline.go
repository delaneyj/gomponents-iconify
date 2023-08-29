package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoDrinksOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.8 2.8l18.4 18.4l-1.425 1.425L13 15.85V19h5v2H6v-2h5v-5l-1.2-1.35l-8.4-8.425L2.8 2.8Zm3.05.2H21v2l-6.2 6.95l-1.425-1.375L14.775 9h-2.95L9.85 7h6.7l1.8-2H7.85l-2-2Zm6.75 6.775Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowReplyAllTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.704 16.294a1 1 0 1 1-1.415 1.414l-4.997-5.004a1 1 0 0 1 0-1.413l4.997-4.998a1 1 0 0 1 1.415 1.414L10.41 11H13a8 8 0 0 1 7.996 7.75L21 19a1 1 0 1 1-2 0a6 6 0 0 0-5.775-5.996L13 13h-2.586l3.29 3.294Zm-5-10.001a1 1 0 0 1 0 1.414l-4.29 4.291l4.29 4.296a1 1 0 1 1-1.415 1.414l-4.997-5.004a1 1 0 0 1 0-1.413L7.29 6.293a1 1 0 0 1 1.415 0Z"/>`),
		g.Group(children),
	)
}
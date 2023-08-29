package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h16v2H3v2h14v4H3v4h14v2H1V3zm18 0h-2v14h2V3zM5 19h16v2H5v-2zM23 7h-2v14h2V7z"/>`),
		g.Group(children),
	)
}
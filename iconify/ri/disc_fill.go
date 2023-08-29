package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 9.17A3 3 0 1 0 15 12V2.458c4.057 1.274 7 5.064 7 9.542c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2c.338 0 .671.017 1 .05v7.12Z"/>`),
		g.Group(children),
	)
}
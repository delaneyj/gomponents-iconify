package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmazonLogoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M152 136a24 24 0 1 1-24-24a24 24 0 0 1 24 24Zm80-8A104 104 0 1 1 128 24a104.11 104.11 0 0 1 104 104Zm-80-24a40 40 0 1 0 0 64a8 8 0 0 0 16 0v-64a40 40 0 0 0-73.87-21.29a8 8 0 0 0 13.54 8.52A24 24 0 0 1 152 104Zm44.81 65.61a8 8 0 0 0-11.2 1.58a72 72 0 0 1-115.22 0a8 8 0 1 0-12.78 9.62a88 88 0 0 0 140.78 0a8 8 0 0 0-1.58-11.2Z"/>`),
		g.Group(children),
	)
}
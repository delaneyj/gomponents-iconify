package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowForwardTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.296 16.294a1 1 0 1 0 1.415 1.414l4.997-5.004a1 1 0 0 0 0-1.413L15.71 6.293a1 1 0 0 0-1.415 1.414L17.59 11H11a8 8 0 0 0-7.996 7.75L3 19a1 1 0 1 0 2 0a6 6 0 0 1 5.775-5.996L11 13h6.586l-3.29 3.294Z"/>`),
		g.Group(children),
	)
}
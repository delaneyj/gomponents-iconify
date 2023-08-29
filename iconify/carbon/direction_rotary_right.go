package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionRotaryRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m22 4l-1.414 1.414L25.172 10H16.91A5.997 5.997 0 1 0 10 16.91V28h2V16.91A6.006 6.006 0 0 0 16.91 12h8.262l-4.586 4.586L22 18l7-7ZM11 15a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}
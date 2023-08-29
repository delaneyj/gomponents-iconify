package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosFastforwardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M48 155l183.5 101L48 356.9V155m224 .8L448 256 272 356.4V156m-16-28v123.2L32 128v256l224-123.2V384l224-128-224-128z" fill="currentColor"/>`),
		g.Group(children),
	)
}
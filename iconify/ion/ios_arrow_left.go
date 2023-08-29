package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M352 115.4L331.3 96 160 256l171.3 160 20.7-19.3L201.5 256z" fill="currentColor"/>`),
		g.Group(children),
	)
}
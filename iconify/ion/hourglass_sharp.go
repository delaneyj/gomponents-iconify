package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M416 32H96v112l108 112L96 368v112h320V368L308 256l108-112ZM272 224v112l91 96H148l92-96V224l-80-80h192Z"/>`),
		g.Group(children),
	)
}
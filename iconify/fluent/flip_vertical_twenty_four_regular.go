package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.656 2.12a.75.75 0 0 1 .344.63v7.5a.75.75 0 0 1-.75.75H2.75a.75.75 0 0 1-.31-1.433l16.5-7.5a.75.75 0 0 1 .716.052ZM6.213 9.5H18.5V3.915L6.213 9.5ZM20 21.5a.5.5 0 0 1-.713.452l-17-8A.5.5 0 0 1 2.5 13h17a.5.5 0 0 1 .5.5v8Z"/>`),
		g.Group(children),
	)
}
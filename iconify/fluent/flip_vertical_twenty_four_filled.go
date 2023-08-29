package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.547 2.163A1 1 0 0 1 20 3v7a1 1 0 0 1-1 1H3a1 1 0 0 1-.4-1.916l16-7a1 1 0 0 1 .947.08ZM7.781 9H18V4.529L7.78 9ZM20 21.5a.5.5 0 0 1-.713.452l-17-8A.5.5 0 0 1 2.5 13h17a.5.5 0 0 1 .5.5v8Z"/>`),
		g.Group(children),
	)
}
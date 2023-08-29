package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserToolTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 2.75a.75.75 0 0 0-1.5 0v14.5A4.75 4.75 0 0 0 7.75 22h8.5A4.75 4.75 0 0 0 21 17.25V2.75a.75.75 0 0 0-1.5 0V7h-15V2.75Zm0 5.75h15V12h-15V8.5Z"/>`),
		g.Group(children),
	)
}
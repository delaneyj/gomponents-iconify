package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullScreenMinimizeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a.5.5 0 0 0 .5.5h1.75a.75.75 0 0 1 0 1.5H12a2 2 0 0 1-2-2V2.25a.75.75 0 0 1 1.5 0V4Zm0 8a.5.5 0 0 1 .5-.5h1.75a.75.75 0 0 0 0-1.5H12a2 2 0 0 0-2 2v1.75a.75.75 0 0 0 1.5 0V12ZM4 11.5a.5.5 0 0 1 .5.5v1.75a.75.75 0 0 0 1.5 0V12a2 2 0 0 0-2-2H2.25a.75.75 0 0 0 0 1.5H4ZM4.5 4a.5.5 0 0 1-.5.5H2.25a.75.75 0 0 0 0 1.5H4a2 2 0 0 0 2-2V2.25a.75.75 0 0 0-1.5 0V4Z"/>`),
		g.Group(children),
	)
}
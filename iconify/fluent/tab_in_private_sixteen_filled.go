package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabInPrivateSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4Zm-.5 2a.5.5 0 0 1 .5-.5h1.793L3.5 5.793V4Zm0 3.207L7.207 3.5h2.586L3.5 9.793V7.207ZM11.207 3.5H12a.5.5 0 0 1 .5.5v.793L4.793 12.5H4a.5.5 0 0 1-.5-.5v-.793L11.207 3.5ZM12.5 6.207v2.586L8.793 12.5H6.207L12.5 6.207Zm0 4V12a.5.5 0 0 1-.5.5h-1.793l2.293-2.293Z"/>`),
		g.Group(children),
	)
}
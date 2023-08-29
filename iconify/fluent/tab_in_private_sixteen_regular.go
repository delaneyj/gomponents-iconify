package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabInPrivateSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 2a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4ZM3 4a1 1 0 0 1 1-1h2.293L3 6.293V4Zm0 3.707L7.707 3h2.586L3 10.293V7.707ZM11.707 3H12a1 1 0 0 1 1 1v.293L4.293 13H4a1 1 0 0 1-1-1v-.293L11.707 3ZM13 5.707v2.586L8.293 13H5.707L13 5.707Zm0 4V12a1 1 0 0 1-1 1H9.707L13 9.707Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabDesktopTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.75 3A3.75 3.75 0 0 0 3 6.75v14.5A3.75 3.75 0 0 0 6.75 25h14.5A3.75 3.75 0 0 0 25 21.25V6.75A3.75 3.75 0 0 0 21.25 3H6.75ZM4.5 6.75A2.25 2.25 0 0 1 6.75 4.5H13v1.75A2.75 2.75 0 0 0 15.75 9h7.75v12.25a2.25 2.25 0 0 1-2.25 2.25H6.75a2.25 2.25 0 0 1-2.25-2.25V6.75Zm19 .75h-7.75c-.69 0-1.25-.56-1.25-1.25V4.5h6.75a2.25 2.25 0 0 1 2.25 2.25v.75Z"/>`),
		g.Group(children),
	)
}
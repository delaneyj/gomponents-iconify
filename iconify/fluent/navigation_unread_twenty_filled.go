package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationUnreadTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.978 4.568a2.25 2.25 0 1 0-.254.764l.004-.01c.125-.23.212-.484.25-.754ZM2.75 4h9.76a3.283 3.283 0 0 0 .24 1.5h-10a.75.75 0 0 1 0-1.5ZM2 9.75A.75.75 0 0 1 2.75 9h14.5a.75.75 0 0 1 0 1.5H2.75A.75.75 0 0 1 2 9.75ZM2.75 14a.75.75 0 0 0 0 1.5h14.5a.75.75 0 0 0 0-1.5H2.75Z"/>`),
		g.Group(children),
	)
}
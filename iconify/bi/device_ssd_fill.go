package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceSsdFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M5 8V4h6v4H5Z"/><path d="M4 0a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2H4Zm0 1.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm9 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM3.5 11a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm9.5-.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM4.75 3h6.5a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-.75.75h-6.5A.75.75 0 0 1 4 8.25v-4.5A.75.75 0 0 1 4.75 3ZM5 12h6a1 1 0 0 1 1 1v2h-1v-2h-.75v2h-1v-2H8.5v2h-1v-2h-.75v2h-1v-2H5v2H4v-2a1 1 0 0 1 1-1Z"/></g>`),
		g.Group(children),
	)
}
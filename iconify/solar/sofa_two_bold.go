package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SofaTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7 13v-1a3 3 0 0 0-2.993-3c.008-.328.027-.566.07-.78A4 4 0 0 1 7.22 5.077C7.606 5 8.07 5 9 5h2.25v8H7Zm5.75 0H17v-1a3 3 0 0 1 2.993-3a4.592 4.592 0 0 0-.07-.78a4 4 0 0 0-3.143-3.143C16.394 5 15.93 5 15 5h-2.25v8Z"/><path d="M18.444 18H5.556a3.57 3.57 0 0 1-.806-.092V19a.75.75 0 0 1-1.5 0v-1.849A3.548 3.548 0 0 1 2 14.444V12a2 2 0 1 1 4 0v1.2a.8.8 0 0 0 .8.8h10.4a.8.8 0 0 0 .8-.8V12a2 2 0 1 1 4 0v2.444a3.548 3.548 0 0 1-1.25 2.707V19a.75.75 0 0 1-1.5 0v-1.092a3.57 3.57 0 0 1-.806.092Z"/></g>`),
		g.Group(children),
	)
}
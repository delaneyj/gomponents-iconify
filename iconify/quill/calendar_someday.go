package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarSomeday(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" stroke="currentColor"><path fill="currentColor" d="M16.75 23.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 17.462a2 2 0 1 1 2 2V20.5M5 12h22m-6-4V4M11 8V4M7 28h18a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2Z"/></g>`),
		g.Group(children),
	)
}
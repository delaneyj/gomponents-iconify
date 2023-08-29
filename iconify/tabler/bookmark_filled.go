package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M15 3a3 3 0 0 1 2.995 2.824L18 6v14a1 1 0 0 1-1.413.911l-.101-.054l-4.487-2.691l-4.485 2.691a1 1 0 0 1-1.508-.743L6 20V6a3 3 0 0 1 2.824-2.995L9 3h6z"/></g>`),
		g.Group(children),
	)
}
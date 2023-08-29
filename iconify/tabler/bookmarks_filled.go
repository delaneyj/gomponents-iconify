package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarksFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M13 6H7a3 3 0 0 0-3 3v12l.006.114a1 1 0 0 0 1.508.743L10 19.166l4.486 2.691A1 1 0 0 0 16 21V9a3 3 0 0 0-3-3z"/><path fill="currentColor" d="M17 2a3 3 0 0 1 2.995 2.824L20 5v12a1 1 0 0 1-1.413.911l-.101-.054l-1-.6a1 1 0 0 1 .489-1.857l.025.001V5a1 1 0 0 0-.883-.993L17 4h-6.001a1 1 0 0 0-.803.401l-.13.197A1 1 0 0 1 8.4 3.5a3 3 0 0 1 2.4-1.494L11 2h6z"/></g>`),
		g.Group(children),
	)
}
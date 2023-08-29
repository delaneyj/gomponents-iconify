package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkBook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-6v8.755a.5.5 0 0 1-.832.374l-1.836-1.632a.5.5 0 0 0-.664 0L7.832 12.13A.5.5 0 0 1 7 11.755V3H5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
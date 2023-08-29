package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashtagLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11.242 3.03a1 1 0 0 1 .728 1.213l-4 16a1 1 0 1 1-1.94-.485l4-16a1 1 0 0 1 1.213-.728zm6 0a1 1 0 0 1 .728 1.213l-4 16a1 1 0 1 1-1.94-.485l4-16a1 1 0 0 1 1.212-.728z"/><path d="M4 9a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1zm-2 6a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1z"/></g>`),
		g.Group(children),
	)
}
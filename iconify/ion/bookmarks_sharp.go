package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarksSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M112 0v48h304v368l48 32V0H112z"/><path fill="currentColor" d="M48 80v432l168-124l168 124V80H48z"/>`),
		g.Group(children),
	)
}
package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExploreSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.25 12a8.75 8.75 0 1 1 17.5 0a8.75 8.75 0 0 1-17.5 0Zm7.6-2.691a5 5 0 0 0-1.756 3.04l-.657 3.949c-.151.911.909 1.523 1.622.936l3.091-2.543a5 5 0 0 0 1.756-3.04l.656-3.949c.152-.911-.908-1.523-1.621-.936L10.85 9.309Z" clip-rule="evenodd"/><path fill="currentColor" d="M10.5 12a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/>`),
		g.Group(children),
	)
}
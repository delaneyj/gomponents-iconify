package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 6h-4V5a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3v1H3a1 1 0 0 0-1 1v4a3 3 0 0 0 1 2.22V19a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5.78A3 3 0 0 0 22 11V7a1 1 0 0 0-1-1ZM9 5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v1H9Zm10 14a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-5h2v1a1 1 0 0 0 2 0v-1h6v1a1 1 0 0 0 2 0v-1h2Zm1-8a1 1 0 0 1-1 1h-2v-1a1 1 0 0 0-2 0v1H9v-1a1 1 0 0 0-2 0v1H5a1 1 0 0 1-1-1V8h16Z"/>`),
		g.Group(children),
	)
}
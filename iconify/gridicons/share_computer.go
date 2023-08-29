package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareComputer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6v2H7v2h10v-2h-3v-2h6a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm0 14H4V4h16zm-3.25-3a1.75 1.75 0 0 1-3.5 0L10 11.36a1.71 1.71 0 1 1 0-2.71L13.25 7a1.77 1.77 0 1 1 .68 1.37L10.71 10l3.22 1.61A1.74 1.74 0 0 1 16.75 13z"/>`),
		g.Group(children),
	)
}
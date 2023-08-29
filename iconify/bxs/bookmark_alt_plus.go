package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkAltPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 2h-12C4.57 2 3 3.57 3 5.5V22l7-3.5l7 3.5v-9h5V5.5C22 3.57 20.43 2 18.5 2zM13 11h-2v2H9v-2H7V9h2V7h2v2h2v2zm7 0h-3V5.5c0-.827.673-1.5 1.5-1.5s1.5.673 1.5 1.5V11z"/>`),
		g.Group(children),
	)
}
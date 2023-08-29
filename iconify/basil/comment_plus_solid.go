package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentPlusSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.823 15.21C2.643 9.857 6.581 4 12.361 4h.321C17 4 20.5 7.5 20.5 11.818a8.732 8.732 0 0 1-8.732 8.732h-7.82a.5.5 0 0 1-.314-.89l1.972-1.583a.5.5 0 0 0 .15-.579l-.933-2.288ZM13.1 9.45a.75.75 0 0 0-1.5 0v2h-2a.75.75 0 0 0 0 1.5h2v2a.75.75 0 0 0 1.5 0v-2h2a.75.75 0 0 0 0-1.5h-2v-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
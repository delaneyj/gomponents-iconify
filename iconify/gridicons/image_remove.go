package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.587 3.423l1.414 1.414L20 6.84V18a2 2 0 0 1-2 2H6.839l-2.006 2.006l-1.414-1.414L20.587 3.423zM12.42 14.42l1.001 1.001l1-1a1.59 1.59 0 0 1 2.11.11L18 16V8.839l-5.58 5.58zM15.162 6H6v6.38l2.19-2.19l1.391 1.391L4 17.162V6a2 2 0 0 1 2-2h11.162l-2 2z"/>`),
		g.Group(children),
	)
}
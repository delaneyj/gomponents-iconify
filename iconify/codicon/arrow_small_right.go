package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSmallRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m8.64 5l2.5 2.5v.7l-2.5 2.5l-.71-.7l1.64-1.65H4v-1h5.57L7.92 5.7l.72-.7z"/>`),
		g.Group(children),
	)
}
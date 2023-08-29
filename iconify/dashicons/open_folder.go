package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenFolder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10.5 6l-2-2H2l.7 9.5L4.2 6h6.3zM5 7L3 17h13l2-10H5z"/>`),
		g.Group(children),
	)
}
package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpImageNotSupported(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.9 21.9l-8.49-8.49L3 3l-.9-.9L.69 3.51L3 5.83V21h15.17l2.31 2.31l1.42-1.41zM5 18l3.5-4.5l2.5 3.01L12.17 15l3 3H5zm16 .17L5.83 3H21v15.17z"/>`),
		g.Group(children),
	)
}
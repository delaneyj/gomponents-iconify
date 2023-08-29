package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpVideocamOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 16.61V6.5l-4 4V6h-6.61zM3.41 1.86L2 3.27L4.73 6H3v12h13.73l3 3l1.41-1.41z"/>`),
		g.Group(children),
	)
}
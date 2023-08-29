package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpFolderOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 6l-2-2H6.83L22 19.17V6zM2.1 2.1L.69 3.51L2 4.83V20h15.17l3.32 3.31l1.41-1.41z"/>`),
		g.Group(children),
	)
}
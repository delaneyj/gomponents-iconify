package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 68H28V48a12 12 0 0 0-24 0v160a12 12 0 0 0 24 0v-28h200v28a12 12 0 0 0 24 0v-96a44.05 44.05 0 0 0-44-44ZM92 156H28V92h64Zm136 0H116V92h92a20 20 0 0 1 20 20Z"/>`),
		g.Group(children),
	)
}
package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.42 4.59l1.167-1.167l1.414 1.414L20 6.84V18a2 2 0 0 1-2 2v-2h-2v2H6.839l-2.006 2.006l-1.414-1.414L4.59 19.42a2.011 2.011 0 0 1-.01-.01L8 16L18 6l1.41-1.42l.01.01zM15.84 11H18V8.839l-2.16 2.16zM16 8.01l.01-.01H16v.011zM6 15.17l-2 2V6a2 2 0 0 1 2-2v2h2V4h9.17l-9 9H6v2.17zM6 8v3h2V8H6zm12 8v-3h-2v3h2z"/>`),
		g.Group(children),
	)
}
package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.59 6.84l5.66 5.66l-5.66 5.66l-.7-.71l4.95-4.95l-4.95-4.95l.7-.71m4 0l5.66 5.66l-5.66 5.66l-.7-.71l4.95-4.95l-4.95-4.95l.7-.71Z"/>`),
		g.Group(children),
	)
}
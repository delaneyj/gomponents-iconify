package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.41 18.16l-5.66-5.66l5.66-5.66l.7.71l-4.95 4.95l4.95 4.95l-.7.71m-4 0L6.75 12.5l5.66-5.66l.7.71l-4.95 4.95l4.95 4.95l-.7.71Z"/>`),
		g.Group(children),
	)
}
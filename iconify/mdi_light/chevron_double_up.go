package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.84 17.41l5.66-5.66l5.66 5.66l-.71.7l-4.95-4.95l-4.95 4.95l-.71-.7m0-4l5.66-5.66l5.66 5.66l-.71.7l-4.95-4.95l-4.95 4.95l-.71-.7Z"/>`),
		g.Group(children),
	)
}
package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.79 4.46l.71-.71L20.25 20.5l-.71.71l-3.2-3.21H4.91l2.33 2.33l-.7.71L3 17.5l3.54-3.54l.7.71L4.91 17h10.43l-9-9H6v4H5V7h.34L2.79 4.46M20 7.5l-3.54 3.54l-.7-.71L18.09 8H9.16l-1-1h9.93l-2.33-2.33l.7-.71L20 7.5M17 13h1v3.84l-1-1V13Z"/>`),
		g.Group(children),
	)
}
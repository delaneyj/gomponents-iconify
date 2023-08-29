package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldMoreVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.17 16.74l-.71-.7L18 12.5l-3.54-3.54l.71-.7l4.24 4.24l-4.24 4.24m-7.34 0L3.59 12.5l4.24-4.24l.71.7L5 12.5l3.54 3.54l-.71.7Z"/>`),
		g.Group(children),
	)
}
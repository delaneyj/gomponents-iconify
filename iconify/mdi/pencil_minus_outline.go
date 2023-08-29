package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PencilMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.1 9l.9.9L5.9 19H5v-.9L14.1 9m3.6-6c-.2 0-.5.1-.7.3l-1.8 1.8l3.7 3.8L20.7 7c.4-.4.4-1 0-1.4l-2.3-2.3c-.2-.2-.5-.3-.7-.3m-3.6 3.2L3 17.2V21h3.8l11-11.1l-3.7-3.7M10 5v2H2V5h8Z"/>`),
		g.Group(children),
	)
}
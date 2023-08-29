package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceHub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-5h4l4-4V8.8q-.9-.325-1.45-1.088T9 6q0-1.25.875-2.125T12 3q1.25 0 2.125.875T15 6q0 .95-.55 1.713T13 8.8V12l4 4h4v5h-5v-3.05l-4-4l-4 4V21H3Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOnOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22V11L6 8V2h12v6l-2 3v11H8Zm4-6.5q-.625 0-1.063-.438T10.5 14q0-.625.438-1.063T12 12.5q.625 0 1.063.438T13.5 14q0 .625-.438 1.063T12 15.5ZM8 5h8V4H8v1Zm8 2H8v.4l2 3V20h4v-9.6l2-3V7Zm-4 5Z"/>`),
		g.Group(children),
	)
}
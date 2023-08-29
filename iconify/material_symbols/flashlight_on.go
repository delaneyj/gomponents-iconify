package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 5V2h12v3H6Zm6 10.5q.625 0 1.063-.438T13.5 14q0-.625-.438-1.063T12 12.5q-.625 0-1.063.438T10.5 14q0 .625.438 1.063T12 15.5ZM8 22V11L6 8V7h12v1l-2 3v11H8Z"/>`),
		g.Group(children),
	)
}
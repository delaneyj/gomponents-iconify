package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SheetsRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 13q-.625 0-1.063-.438T4 11.5v-9q0-.625.438-1.063T5.5 1h13q.625 0 1.063.438T20 2.5v9q0 .625-.438 1.063T18.5 13h-13ZM6 6h5V3H6v3Zm7 0h5V3h-5v3Zm-2 5V8H6v3h5Zm2 0h5V8h-5v3ZM7 22l-4-4l4-4l1.4 1.4L6.825 17H20v2H6.825L8.4 20.6L7 22Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothConnected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22v-7.6L6.4 19L5 17.6l5.6-5.6L5 6.4L6.4 5L11 9.6V2h1l5.7 5.7l-4.3 4.3l4.3 4.3L12 22h-1Zm2-12.4l1.9-1.9L13 5.85V9.6Zm0 8.55l1.9-1.85l-1.9-1.9v3.75ZM5 13.5q-.625 0-1.063-.438T3.5 12q0-.625.438-1.063T5 10.5q.625 0 1.063.438T6.5 12q0 .625-.438 1.063T5 13.5Zm14 0q-.625 0-1.063-.438T17.5 12q0-.625.438-1.063T19 10.5q.625 0 1.063.438T20.5 12q0 .625-.438 1.063T19 13.5Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothDriveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 16q.625 0 1.063-.438T8 14.5q0-.625-.438-1.063T6.5 13q-.625 0-1.063.438T5 14.5q0 .625.438 1.063T6.5 16Zm9 0q.625 0 1.063-.438T17 14.5q0-.625-.438-1.063T15.5 13q-.625 0-1.063.438T14 14.5q0 .625.438 1.063T15.5 16ZM3 21q-.425 0-.713-.288T2 20v-8l2.1-6q.15-.45.537-.725T5.5 5H15v2H5.85L4.8 10H15v2H4v5h14v-5h2v8q0 .425-.288.713T19 21h-1q-.425 0-.713-.288T17 20v-1H5v1q0 .425-.288.713T4 21H3Zm15.65-10V7.2l-2.3 2.3l-.7-.7l2.8-2.8l-2.8-2.8l.7-.7l2.3 2.3V1h.5L22 3.9L19.85 6L22 8.15L19.15 11h-.5Zm1-6.2l.95-.9l-.95-.95V4.8Zm0 4.3l.95-.95l-.95-.95v1.9ZM4 12v5v-5Z"/>`),
		g.Group(children),
	)
}
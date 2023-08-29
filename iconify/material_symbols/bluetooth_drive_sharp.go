package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothDriveSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 16q.625 0 1.063-.438T8 14.5q0-.625-.438-1.063T6.5 13q-.625 0-1.063.438T5 14.5q0 .625.438 1.063T6.5 16Zm9 0q.625 0 1.063-.438T17 14.5q0-.625-.438-1.063T15.5 13q-.625 0-1.063.438T14 14.5q0 .625.438 1.063T15.5 16ZM2 21v-9l2.45-7H15v2H5.85L4.8 10H15v2h5v9h-3v-2H5v2H2Zm16.65-10V7.2l-2.3 2.3l-.7-.7l2.8-2.8l-2.8-2.8l.7-.7l2.3 2.3V1h.5L22 3.9L19.85 6L22 8.15L19.15 11h-.5Zm1-6.2l.95-.9l-.95-.95V4.8Zm0 4.3l.95-.95l-.95-.95v1.9Z"/>`),
		g.Group(children),
	)
}
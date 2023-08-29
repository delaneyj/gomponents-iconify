package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MixtureMed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22.5L6 21v-4H5q-.825 0-1.413-.588T3 15V7.5q-.425 0-.713-.288T2 6.5q0-.425.288-.713T3 5.5h3V4h-.5q-.425 0-.713-.288T4.5 3q0-.425.288-.713T5.5 2h3q.425 0 .713.288T9.5 3q0 .425-.288.713T8.5 4H8v1.5h3q.425 0 .713.288T12 6.5q0 .425-.288.713T11 7.5V15q0 .825-.588 1.413T9 17H8v5.5Zm9-.5q-1.65 0-2.825-1.175T13 18v-8q0-1.65 1.175-2.825T17 6q1.65 0 2.825 1.175T21 10v8q0 1.65-1.175 2.825T17 22ZM5 15h4v-1.5H7.25q-.3 0-.525-.225T6.5 12.75q0-.3.225-.525T7.25 12H9v-1.5H7.25q-.3 0-.525-.225T6.5 9.75q0-.3.225-.525T7.25 9H9V7.5H5V15Zm10 1h4v-4h-4v4Z"/>`),
		g.Group(children),
	)
}
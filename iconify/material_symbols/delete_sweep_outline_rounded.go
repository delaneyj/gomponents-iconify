package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteSweepOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19q-.825 0-1.413-.588T3 17V8q-.425 0-.713-.288T2 7q0-.425.288-.713T3 6h3v-.5q0-.425.288-.713T7 4.5h2q.425 0 .713.288T10 5.5V6h3q.425 0 .713.288T14 7q0 .425-.288.713T13 8v9q0 .825-.588 1.413T11 19H5Zm11-1q-.425 0-.713-.288T15 17q0-.425.288-.713T16 16h2q.425 0 .713.288T19 17q0 .425-.288.713T18 18h-2Zm0-4q-.425 0-.713-.288T15 13q0-.425.288-.713T16 12h4q.425 0 .713.288T21 13q0 .425-.288.713T20 14h-4Zm0-4q-.425 0-.713-.288T15 9q0-.425.288-.713T16 8h5q.425 0 .713.288T22 9q0 .425-.288.713T21 10h-5ZM5 8v9h6V8H5Z"/>`),
		g.Group(children),
	)
}
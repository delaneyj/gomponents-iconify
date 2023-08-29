package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdCardRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 11q.425 0 .713-.288T11 10V8q0-.425-.288-.713T10 7q-.425 0-.713.288T9 8v2q0 .425.288.713T10 11Zm3 0q.425 0 .713-.288T14 10V8q0-.425-.288-.713T13 7q-.425 0-.713.288T12 8v2q0 .425.288.713T13 11Zm3 0q.425 0 .713-.288T17 10V8q0-.425-.288-.713T16 7q-.425 0-.713.288T15 8v2q0 .425.288.713T16 11ZM6 22q-.825 0-1.413-.588T4 20V8.825q0-.4.15-.762t.425-.638l4.85-4.85q.275-.275.638-.425t.762-.15H18q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22H6Z"/>`),
		g.Group(children),
	)
}
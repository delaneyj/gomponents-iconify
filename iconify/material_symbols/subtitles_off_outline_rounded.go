package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtitlesOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4l2 2H4v12h11.15l-2-2H7q-.425 0-.713-.288T6 15q0-.425.288-.713T7 14h4.15l-9.8-9.85q-.275-.275-.288-.688t.288-.712q.275-.275.7-.275t.7.275l18.5 18.5q.275.275.288.688t-.288.712q-.275.275-.7.275t-.7-.275L17.15 20H4Zm17.75-1.1L20 17.15V6H8.85l-2-2H20q.825 0 1.413.588T22 6v11.9q0 .275-.05.525t-.2.475Zm-6.9-6.9l-2-2H17q.425 0 .713.288T18 11q0 .425-.288.713T17 12h-2.15Zm-.425-.425Zm-4.85.85ZM7 12q-.425 0-.713-.288T6 11q0-.425.288-.713T7 10q.425 0 .713.288T8 11q0 .425-.288.713T7 12Z"/>`),
		g.Group(children),
	)
}
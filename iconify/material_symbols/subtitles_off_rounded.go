package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtitlesOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 12q.425 0 .713-.288T8 11q0-.425-.288-.713T7 10q-.425 0-.713.288T6 11q0 .425.288.713T7 12Zm-.15-8H20q.825 0 1.413.588T22 6v11.9q0 .275-.05.525t-.2.475l-6.9-6.9H17q.425 0 .713-.288T18 11q0-.425-.288-.713T17 10h-4.15l-6-6ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4l8 8H9.2L1.35 4.15q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l18.5 18.5q.275.275.288.688t-.288.712q-.275.275-.7.275t-.7-.275L17.15 20H4Zm7.15-6H7q-.425 0-.713.288T6 15q0 .425.288.713T7 16h6.15l-2-2Z"/>`),
		g.Group(children),
	)
}
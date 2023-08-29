package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReportOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.5 23.3l-3.55-3.55l-1.2 1.25h-7.5L3 15.7V8.25l1.2-1.2L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4Zm-.7-6.35L13 10.2V7h-2v1.2l-3.95-4L8.25 3h7.5L21 8.25v7.45l-1.2 1.25ZM12 17q.425 0 .713-.288T13 16v-.15l-.85-.85H12q-.425 0-.713.288T11 16q0 .425.288.713T12 17Z"/>`),
		g.Group(children),
	)
}
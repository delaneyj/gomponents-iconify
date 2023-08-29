package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRollOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 20h-8q0 .825-.588 1.413T12 22H4q-.825 0-1.413-.588T2 20V5q0-.825.588-1.413T4 3h1V2q0-.425.288-.713T6 1h4q.425 0 .713.288T11 2v1h1q.825 0 1.413.588T14 5h8v15Zm-2-2V7h-8V5H9V3H7v2H4v15h8v-2h8ZM9 17h2v-2H9v2Zm0-7h2V8H9v2Zm4 7h2v-2h-2v2Zm0-7h2V8h-2v2Zm4 7h2v-2h-2v2Zm0-7h2V8h-2v2Zm-5 1.5Z"/>`),
		g.Group(children),
	)
}
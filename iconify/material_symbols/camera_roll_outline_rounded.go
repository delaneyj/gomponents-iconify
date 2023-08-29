package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRollOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V5q0-.825.588-1.413T4 3h1V2q0-.425.288-.713T6 1h4q.425 0 .713.288T11 2v1h1q.825 0 1.413.588T14 5h6q.825 0 1.413.588T22 7v11q0 .825-.588 1.413T20 20h-6q0 .825-.588 1.413T12 22H4Zm0-2h8v-2h8V7h-8V5H4v15Zm5.5-3h1q.2 0 .35-.15t.15-.35v-1q0-.2-.15-.35T10.5 15h-1q-.2 0-.35.15T9 15.5v1q0 .2.15.35t.35.15Zm0-7h1q.2 0 .35-.15T11 9.5v-1q0-.2-.15-.35T10.5 8h-1q-.2 0-.35.15T9 8.5v1q0 .2.15.35t.35.15Zm4 7h1q.2 0 .35-.15t.15-.35v-1q0-.2-.15-.35T14.5 15h-1q-.2 0-.35.15t-.15.35v1q0 .2.15.35t.35.15Zm0-7h1q.2 0 .35-.15T15 9.5v-1q0-.2-.15-.35T14.5 8h-1q-.2 0-.35.15T13 8.5v1q0 .2.15.35t.35.15Zm4 7h1q.2 0 .35-.15t.15-.35v-1q0-.2-.15-.35T18.5 15h-1q-.2 0-.35.15t-.15.35v1q0 .2.15.35t.35.15Zm0-7h1q.2 0 .35-.15T19 9.5v-1q0-.2-.15-.35T18.5 8h-1q-.2 0-.35.15T17 8.5v1q0 .2.15.35t.35.15ZM4 20V5v15Z"/>`),
		g.Group(children),
	)
}
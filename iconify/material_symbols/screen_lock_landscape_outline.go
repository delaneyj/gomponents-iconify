package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenLockLandscapeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.85 16q-.35 0-.6-.25t-.25-.6v-3.3q0-.35.25-.6t.6-.25H10v-1q0-.825.588-1.412T12 8q.825 0 1.413.588T14 10v1h.15q.35 0 .6.25t.25.6v3.3q0 .35-.25.6t-.6.25h-4.3ZM11 11h2v-1q0-.425-.287-.713T12 9q-.425 0-.713.288T11 10v1Zm-8 8q-.825 0-1.413-.588T1 17V7q0-.825.588-1.413T3 5h18q.825 0 1.413.588T23 7v10q0 .825-.588 1.413T21 19H3Zm1-2V7H3v10h1Zm2 0h12V7H6v10Zm14 0h1V7h-1v10ZM4 7H3h1Zm16 0h1h-1Z"/>`),
		g.Group(children),
	)
}
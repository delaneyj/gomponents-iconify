package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenLockRotationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.95 11.65l1.4-1.4l1.1 1.1q.575.575.575 1.35t-.575 1.35L16.1 18.4q-.575.575-1.35.575T13.4 18.4l-7.85-7.85q-.575-.575-.575-1.35t.575-1.35L9.9 3.5q.575-.575 1.35-.575t1.35.575l1.1 1.1l-1.35 1.45l-1.1-1.1L7 9.2l7.75 7.75L19 12.7l-1.05-1.05ZM13 22.95q-2.475 0-4.663-.95t-3.812-2.575Q2.9 17.8 1.95 15.612T1 10.95h2q0 1.8.613 3.425T5.3 17.3q1.075 1.3 2.55 2.212t3.175 1.263L8.4 18.15l1.4-1.4l5.9 5.9q-.675.15-1.325.225T13 22.95ZM15.85 9q-.35 0-.6-.25t-.25-.6v-3.3q0-.35.25-.6t.6-.25H16V3q0-.825.588-1.413T18 1q.825 0 1.413.588T20 3v1h.15q.35 0 .6.25t.25.6v3.3q0 .35-.25.6t-.6.25h-4.3ZM17 4h2V3q0-.425-.288-.712T18 2q-.425 0-.713.288T17 3v1Zm-4 6.95Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamStandSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.8 21l.8-9.2q.125-1.8 1.05-3.275T8 6.15V10q0 1.65 1.175 2.825T12 14q1.675 0 2.838-1.175T16 10V6.15q1.425.9 2.35 2.375T19.4 11.8l.775 9.2H3.8Zm8.2-9q-.825 0-1.412-.588T10 10V6q0-.825.588-1.413T12 4q.825 0 1.413.588T14 6v4q0 .825-.588 1.413T12 12Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoPhotographyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 19.125l-2-2V7h-4.05l-1.825-2h-4.25l-.95 1.05L7.5 4.625L9 3h6l1.85 2H20q.825 0 1.413.588T22 7v12.125ZM4 21q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h1.025l2 2H4v12h15.025l2 2H4Zm11.65-5.375q-.625.85-1.563 1.363T12 17.5q-1.875 0-3.188-1.313T7.5 13q0-1.15.513-2.087T9.374 9.35l1.45 1.45q-.6.325-.963.9T9.5 13q0 1.05.725 1.775T12 15.5q.725 0 1.3-.363t.9-.962l1.45 1.45ZM15.2 9.8q.625.6.963 1.425T16.5 13v.3q0 .15-.025.3L11.4 8.525q.15-.025.3-.025h.3q.95 0 1.775.338T15.2 9.8Zm5.275 13.5L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM10.175 13Zm4.275-1.425Z"/>`),
		g.Group(children),
	)
}
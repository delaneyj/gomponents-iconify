package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoPhotographyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 19.125l-2-2V7h-4.05l-1.825-2h-4.25l-.95 1.05L7.5 4.625l.9-.975q.275-.3.663-.475T9.874 3h4.25q.425 0 .813.175t.662.475L16.85 5H20q.825 0 1.413.588T22 7v12.125ZM4 21q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h1.025l2 2H4v12h15.025l2 2H4Zm11.65-5.375q-.625.85-1.563 1.363T12 17.5q-1.875 0-3.188-1.313T7.5 13q0-1.15.513-2.087T9.374 9.35l1.45 1.45q-.6.325-.963.9T9.5 13q0 1.05.725 1.775T12 15.5q.725 0 1.3-.363t.9-.962l1.45 1.45ZM15.2 9.8q.625.6.963 1.425T16.5 13v.3q0 .15-.025.3L11.4 8.525q.15-.025.3-.025h.3q.95 0 1.775.338T15.2 9.8Zm4.575 12.775L1.375 4.2q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3ZM10.175 13Zm4.275-1.425Z"/>`),
		g.Group(children),
	)
}
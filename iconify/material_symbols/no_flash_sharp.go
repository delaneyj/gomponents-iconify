package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoFlashSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.475 23.3l-2.5-2.5V22H2V9.4h4.15l.2-.225L.675 3.5L2.1 2.075l19.8 19.8l-1.425 1.425ZM9.25 12.075q-1.425.3-2.337 1.375T6 16q0 1.65 1.175 2.825T10 20q1.475 0 2.575-.925t1.375-2.325l-.013.025l.013-.025l-4.7-4.675ZM10 18q-.825 0-1.413-.588T8 16q0-.825.588-1.413T10 14q.825 0 1.413.588T12 16q0 .825-.588 1.413T10 18Zm8-2.875L10.875 8h1.7l1.275 1.4H18v5.725ZM19 11V7h-1V2h4l-1.6 3.6H22L19 11Z"/>`),
		g.Group(children),
	)
}
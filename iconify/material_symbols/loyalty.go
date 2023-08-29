package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loyalty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 17.5l3.5-3.5q.275-.275.438-.65t.162-.8q0-.85-.6-1.45t-1.45-.6q-.475 0-.938.275T13 11.7q-.75-.7-1.175-.95t-.875-.25q-.85 0-1.45.6t-.6 1.45q0 .425.163.8T9.5 14l3.5 3.5Zm1.25 3.9q-.575.575-1.425.575T11.4 21.4l-8.8-8.8q-.275-.275-.438-.65T2 11.15V4q0-.825.588-1.413T4 2h7.15q.425 0 .8.163t.65.437l8.8 8.825q.575.575.575 1.413T21.4 14.25l-7.15 7.15ZM6.5 8q.625 0 1.063-.438T8 6.5q0-.625-.438-1.063T6.5 5q-.625 0-1.063.438T5 6.5q0 .625.438 1.063T6.5 8Z"/>`),
		g.Group(children),
	)
}
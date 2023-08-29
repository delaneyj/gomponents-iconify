package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M4.318 20.535C5.636 22 7.758 22 12 22c4.243 0 6.364 0 7.682-1.465C21 19.072 21 16.714 21 12s0-7.071-1.318-8.536C18.364 2 16.242 2 12 2C7.757 2 5.636 2 4.318 3.464C3 4.93 3 7.286 3 12c0 4.714 0 7.071 1.318 8.535Z"/><path stroke="currentColor" stroke-width="1.5" d="M7 8c0-.465 0-.697.051-.888a1.5 1.5 0 0 1 1.06-1.06C8.304 6 8.536 6 9 6h6c.465 0 .697 0 .888.051a1.5 1.5 0 0 1 1.06 1.06C17 7.304 17 7.536 17 8s0 .697-.051.888a1.5 1.5 0 0 1-1.06 1.06C15.697 10 15.464 10 15 10H9c-.465 0-.697 0-.888-.051a1.5 1.5 0 0 1-1.06-1.06C7 8.696 7 8.464 7 8Z" opacity=".5"/><circle cx="8" cy="13" r="1" fill="currentColor" opacity=".5"/><circle cx="8" cy="17" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="13" r="1" fill="currentColor" opacity=".5"/><circle cx="12" cy="17" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="13" r="1" fill="currentColor" opacity=".5"/><circle cx="16" cy="17" r="1" fill="currentColor" opacity=".5"/></g>`),
		g.Group(children),
	)
}
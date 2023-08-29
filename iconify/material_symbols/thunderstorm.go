package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thunderstorm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.55 24l2.35-2.675l-2-1L14.8 17h2.65l-2.35 2.675l2 1L14.2 24h-2.65Zm-6 0l2.35-2.675l-2-1L8.8 17h2.65L9.1 19.675l2 1L8.2 24H5.55Zm1.95-8q-2.275 0-3.887-1.613T2 10.5q0-2.075 1.375-3.625t3.4-1.825q.8-1.425 2.188-2.238T12 2q2.25 0 3.913 1.438t2.012 3.587q1.725.15 2.9 1.425T22 11.5q0 1.875-1.312 3.188T17.5 16h-10Z"/>`),
		g.Group(children),
	)
}
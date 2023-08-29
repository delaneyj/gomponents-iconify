package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarbudsBatterySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 18V7h2V6h2v1h2v11h-6ZM5.375 18q-1.425 0-2.4-.975T2 14.625V6h4v4H3.5v4.625q0 .8.537 1.338t1.338.537q.8 0 1.338-.537t.537-1.338v-5.25q0-1.425.975-2.4t2.4-.975q1.425 0 2.4.975t.975 2.4V18h-4v-4h2.5V9.375q0-.8-.537-1.338T10.625 7.5q-.8 0-1.338.537T8.75 9.375v5.25q0 1.425-.975 2.4t-2.4.975Z"/>`),
		g.Group(children),
	)
}
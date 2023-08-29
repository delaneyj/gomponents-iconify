package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polycam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.044 19.746l-7.425-5.628L11.127 24l-1.7-1.337v-9.618l1.7 1.307L23.619 4.5l13.836 10.297"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.648 29.155l6.983-4.99l-.176-9.368l-12.924 9.972"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.53 24.77l12.384 8.877l1.658-1.217v9.853L36.914 43.5L24.53 34.153v-9.384Z"/>`),
		g.Group(children),
	)
}
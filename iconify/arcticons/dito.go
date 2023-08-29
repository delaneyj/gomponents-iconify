package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dito(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.5 5.5h-17a8 8 0 0 0-8 8v29h25c6.627 0 12-5.373 12-12v-13c0-6.627-5.373-12-12-12Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.168 36.682l2.102-4.832V14.257h11.079c5.016 0 8.96 4.192 8.96 9.585c0 5.625-3.944 9.584-8.96 9.584h-7.351l-5.83 3.256Z"/>`),
		g.Group(children),
	)
}
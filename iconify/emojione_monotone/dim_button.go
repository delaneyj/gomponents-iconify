package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DimButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M29 51.098h6V58h-6zM29 6h6v6.902h-6zm22.098 23.001H58v6h-6.902zM6 29.001h6.902v6H6zm9.736 23.505l-4.242-4.242l4.881-4.88l4.242 4.241zm32.528-41.012l4.242 4.241l-4.88 4.88l-4.243-4.241zm-.003 41.013l-4.88-4.88l4.24-4.242l4.882 4.88zM15.736 11.494l4.881 4.88l-4.242 4.243l-4.88-4.881zM32 17c-8.285 0-15 6.716-15 15s6.715 15 15 15c8.283 0 15-6.716 15-15s-6.717-15-15-15m0 25c-5.523 0-10-4.478-10-10s4.477-10 10-10c5.521 0 10 4.478 10 10s-4.479 10-10 10"/>`),
		g.Group(children),
	)
}
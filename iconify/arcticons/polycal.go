package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polycal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.353 31.024s-1.58 3.334.015 4.001c2.62 1.09 3.498-8.915 2.41-17.433M33.616 30.79s-.16 4.68-4.048 4.154c-4.072-.555-3.58-10.917-2.492-16.849"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.556 20.728s2.188-3.089 7.792-3.627c5.99-.574 12.987 2.89 13.69.39c.44-1.79-1.58-2.234-3.663-1.055"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.1 7.86a1.92 1.92 0 0 0-1.92 1.92v31.8A1.92 1.92 0 0 0 8.1 43.5h31.8a1.92 1.92 0 0 0 1.92-1.92V9.78a1.92 1.92 0 0 0-1.92-1.92Zm4.7 2.44V4.5m22.4 5.8V4.5m-16.8 5.8V4.5m5.6 5.8V4.5m5.6 5.8V4.5"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pixeldroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.421 14.235h-8.244a1.818 1.818 0 0 0-1.818 1.819V34.66a.574.574 0 0 0 1 .385l5.325-5.885a1.148 1.148 0 0 1 .852-.378h3.127a7.273 7.273 0 0 0 7.271-7.41a7.432 7.432 0 0 0-7.513-7.137Z"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.548 12.958a6.973 6.973 0 0 1 5.149-2.261a13.256 13.256 0 0 1 9.035 3.538m9.061-11.199a6.972 6.972 0 0 1 3.748 4.2a13.254 13.254 0 0 1-.48 9.482m13.356 5.379a6.973 6.973 0 0 1-2.835 4.846A13.299 13.299 0 0 1 29.268 28.3m3.181 15.477a6.973 6.973 0 0 1-5.506-1.195a13.298 13.298 0 0 1-5.414-12.144M7.79 38.125a6.972 6.972 0 0 1-.553-5.584a13.292 13.292 0 0 1 9.122-8.717"/>`),
		g.Group(children),
	)
}
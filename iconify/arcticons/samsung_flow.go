package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungFlow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.077 24.234H14.114c-5.042 0-9.364 4.045-9.24 9.086a8.864 8.864 0 0 0 8.639 8.638c5.04.124 9.085-4.197 9.085-9.24V15.143a6.099 6.099 0 0 1 6.1-6.099h5.379"/><circle cx="37.974" cy="24.234" r="3.897" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="37.974" cy="9.044" r="3.897" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
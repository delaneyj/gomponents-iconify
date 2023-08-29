package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anonaddy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.74 37.21h30.59A3.17 3.17 0 0 0 43.5 34V12.91a2.12 2.12 0 0 0-2.11-2.12H10.8a1.06 1.06 0 0 0-.63 1.91L25.54 24l16.88-12.41M11.86 29.25l7.22-5.99m-9.14 7.58l-1.46 1.21m.34-9.43l4.54-3.71m-6.46 5.3l-1.46 1.21m10.81-4.39l-1.74 1.44m-1.96 1.63L4.5 30.78"/>`),
		g.Group(children),
	)
}
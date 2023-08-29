package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirNz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.375 26.972H6.592S9.382 42.5 18.48 42.5c7.764 0 6.066-9.058 1.577-9.058c-3.477 0-3.073 3.114-1.092 2.548m22.443-14.962H9.625S19.33 5.5 28.428 5.5c7.764 0 6.065 9.058 1.577 9.058c-3.478 0-3.073-3.114-1.092-2.548"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Virtuosity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.38 4.5s4.36 7.5 4.36 10.05s-5.08 5.36-5.08 8.07c0 3 8.34 12.87 8.34 12.87s-10-6.63-10 .61a8.78 8.78 0 0 0 3.76 7.4"/>`),
		g.Group(children),
	)
}
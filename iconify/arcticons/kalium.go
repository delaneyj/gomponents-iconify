package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kalium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37 21.3c-1.187 3.442-4.349 7.197-7.919 8.053c-3.661.878-7.591-3.116-10.117-5.054M11 26.7c1.187-3.443 4.349-7.197 7.919-8.053c3.66-.878 7.591 3.116 10.117 5.053"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
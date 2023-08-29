package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Accrescent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.745 2.501a18.517 18.517 0 0 1 7.89 10.791c2.647 9.879-3.216 20.033-13.094 22.68a18.518 18.518 0 0 1-13.31-1.497A21.504 21.504 0 0 0 29.556 44.77c11.472-3.074 18.28-14.865 15.206-26.337A21.504 21.504 0 0 0 23.745 2.502Z"/>`),
		g.Group(children),
	)
}
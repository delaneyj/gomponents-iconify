package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doctolib(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.888 11.548a120.35 120.35 0 0 0-4.89 20.964"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.288 6.027c9.294-.908 20.824-1.537 24.108 6.08s1.048 19.077-2.166 23.62s-10.46 7.448-17.4 6.638c-9.573-1.118-16.845-4.659-16.841-12.858a8.636 8.636 0 0 1 1.892-5.062"/>`),
		g.Group(children),
	)
}
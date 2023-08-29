package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tec(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 24h39m-25.907 7.564h7.564m-3.501-15.128h7.564m-5.18 5.66h2.795m-5.179-5.66l-4.063 15.128m-9.24-15.128h10.023M9.302 31.564l4.063-15.128m25.487 10.945a7.033 7.033 0 0 1-6.067 4.184c-2.767 0-4.408-2.244-3.665-5.012l1.37-5.105a7.016 7.016 0 0 1 6.358-5.012a3.625 3.625 0 0 1 3.835 4.002"/>`),
		g.Group(children),
	)
}
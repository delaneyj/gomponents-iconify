package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rupee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 12a2 2 0 0 1 2-2h12a2 2 0 1 1 0 4h-3.07a7.96 7.96 0 0 1 .818 2H30a2 2 0 1 1 0 4h-2.252a8.012 8.012 0 0 1-6.121 5.834l6.037 9.057a2 2 0 1 1-3.328 2.218l-8-12A2 2 0 0 1 18 22h2c1.48 0 2.773-.804 3.465-2H18a2 2 0 1 1 0-4h5.465A3.998 3.998 0 0 0 20 14h-2a2 2 0 0 1-2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
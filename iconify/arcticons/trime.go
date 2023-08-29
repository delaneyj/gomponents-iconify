package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 41.41h39M24 6.59v34.82m14.58-31.12V27a2.11 2.11 0 0 1-2.11 2.11H11.53A2.11 2.11 0 0 1 9.42 27V10.29"/>`),
		g.Group(children),
	)
}
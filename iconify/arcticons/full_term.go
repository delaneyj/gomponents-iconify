package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullTerm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 33.754c-5.361 4.24-12.13 6.77-19.5 6.77s-14.139-2.53-19.5-6.77L17.175 7.477c2.064 1 4.377 1.56 6.825 1.56s4.76-.56 6.825-1.56L43.5 33.754Z"/>`),
		g.Group(children),
	)
}
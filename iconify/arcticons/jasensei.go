package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jasensei(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.173 13.827v15.26a5.102 5.102 0 0 1-5.086 5.086h0A5.102 5.102 0 0 1 11 29.087v-1.78m26 6.866l-6.613-20.346l-6.867 20.346m2.289-6.867h8.902"/>`),
		g.Group(children),
	)
}
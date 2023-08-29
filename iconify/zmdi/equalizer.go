package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 363V21h85v342h-85zM0 363V192h85v171H0zm256-235h85v235h-85V128z"/>`),
		g.Group(children),
	)
}
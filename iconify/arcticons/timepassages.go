package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timepassages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.808 7a16.996 16.996 0 0 0-4.603.648a16.988 16.988 0 0 1-.013 32.7a16.96 16.96 0 0 0 21.616-16.346V24A17 17 0 0 0 23.81 7Z"/>`),
		g.Group(children),
	)
}
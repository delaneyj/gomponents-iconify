package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Debiandroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.835 15.125a5.08 5.08 0 0 0-3.246-1.358c-3.467 0-6.533 3.666-6.533 6.8a7.603 7.603 0 0 0 7.933 7.733a9.741 9.741 0 0 0 9.733-9.933c0-6.334-2.496-13.867-12.733-13.867c-10.467 0-16.711 8.267-16.711 17.6c0 11.6 7.444 21.4 16.978 21.4"/>`),
		g.Group(children),
	)
}
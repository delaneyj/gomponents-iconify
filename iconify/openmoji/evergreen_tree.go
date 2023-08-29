package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EvergreenTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5C9E31" d="m56.035 54.005l-8.123-15.21l3.346.54a.998.998 0 0 0 1.018-1.496l-7.372-12.434l3.077.218a.998.998 0 0 0 .89-1.565L36.825 6.724a1.019 1.019 0 0 0-.834-.428a.998.998 0 0 0-.822.453l-3.717 5.7l-7.821 12.008a.998.998 0 0 0 .998 1.53l2.301-.374l-7.552 12.212a.998.998 0 0 0 .998 1.512l3.727-.563l-8.135 15.231a.998.998 0 0 0 .947 1.466l12.447-.836l1.036-.07l10.961-.125l13.718 1.031h.075a.998.998 0 0 0 .88-1.468l.003.002z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="m41.448 53.49l13.746 1.033l-9.09-17.02l5.344.862l-8.359-14.099l4.984.349L36.004 7.246L24.448 24.99l4.47-.724l-8.72 14.099l5.705-.862l-9.09 17.02l13.51-.908"/><path stroke-miterlimit="10" d="M36.004 55.964v9.309"/></g>`),
		g.Group(children),
	)
}
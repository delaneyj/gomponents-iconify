package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Universal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="4" d="M24 38c7.732 0 14-6.268 14-14s-6.268-14-14-14s-14 6.268-14 14s6.268 14 14 14Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M11 29c1.509.624 4 1 5.259-.468c1.258-1.469.136-3.78 1.53-4.564c1.528-.86 2.631 2.064 5.502 1.548S28 21 28 19s-1.715-2-1.838-3.946C26 12.5 28 11 28 11m0 26c-1.086-.909-2-1.5-2-3s1-1 2-2s.5-3 1.5-3.5s4.108.556 6.5 2.5"/><circle cx="24" cy="4" r="2" fill="currentColor"/><circle cx="24" cy="44" r="2" fill="currentColor"/><circle cx="44" cy="24" r="2" fill="currentColor"/><circle cx="38" cy="10" r="2" fill="currentColor"/><circle cx="10" cy="38" r="2" fill="currentColor"/><circle cx="4" cy="24" r="2" fill="currentColor"/><circle cx="10" cy="10" r="2" fill="currentColor"/><circle cx="38" cy="38" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 24c0 3.815 1.526 7.273 4 9.798M24 38c7.732 0 14-6.268 14-14M24 10c3.815 0 7.273 1.526 9.798 4"/></g>`),
		g.Group(children),
	)
}
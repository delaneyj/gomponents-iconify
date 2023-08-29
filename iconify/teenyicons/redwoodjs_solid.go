package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedwoodjsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.252.066a.5.5 0 0 1 .496 0l3.36 1.92L7.5 4.337L3.893 1.985L7.252.065ZM3.185 2.718L1.96 4.25L4.2 6.49l2.386-1.556l-3.4-2.217ZM1.438 5.146L.52 8.363a.5.5 0 0 0 .034.36l.054.11l2.735-1.784l-1.904-1.903Zm-.38 4.587l1.371 2.743l4.227-2.113l-2.591-2.591l-3.007 1.961Zm2.219 3.437l4.02 1.787a.5.5 0 0 0 .406 0l4.02-1.787L7.5 11.06l-4.223 2.11Zm9.293-.694l1.372-2.743l-3.007-1.961l-2.59 2.591l4.226 2.113Zm1.823-3.643l.054-.11a.5.5 0 0 0 .034-.36l-.92-3.217l-1.903 1.903l2.735 1.784Zm-1.352-4.582l-1.226-1.533l-3.4 2.217l2.387 1.556l2.24-2.24ZM7.5 5.532l2.58 1.682L7.5 9.793l-2.58-2.58l2.58-1.68Z"/>`),
		g.Group(children),
	)
}
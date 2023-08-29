package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightTakeoffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.949 10.112a1.5 1.5 0 0 1-1.06 1.837L5.221 16.147a1 1 0 0 1-1.133-.48l-2.623-4.725l1.449-.389l2.468 2.445l5.095-1.365l-4.51-7.074l1.931-.518l6.952 6.42l5.26-1.41a1.5 1.5 0 0 1 1.838 1.061ZM4 19h16v2H4v-2Z"/>`),
		g.Group(children),
	)
}
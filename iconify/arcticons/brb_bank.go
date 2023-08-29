package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrbBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.356 39.415h27.397L42.5 8.585H9.443L5.558 33.82a4.855 4.855 0 0 0 4.798 5.594Zm6.178-9.758l16.187-13.321"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.111 39.415c0-13.653-10.139-24.938-23.298-26.735"/>`),
		g.Group(children),
	)
}
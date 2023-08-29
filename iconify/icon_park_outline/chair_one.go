package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChairOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M29.454 23H18.545C15.819 23 14 21.333 14 18.833V12h4V6a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v6h4v6.833c0 2.5-1.818 4.167-4.546 4.167ZM20 23l-8 21m16-21l8 21m-6-32h8m-20 0h-8m7 19h14m-16 7h18"/>`),
		g.Group(children),
	)
}
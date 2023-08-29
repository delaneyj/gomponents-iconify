package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoveeHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.646 14.626C8.129 7.448 15.486 2.5 24 2.5a21.433 21.433 0 0 1 15.203 6.297M45.5 24c0 11.874-9.626 21.5-21.5 21.5c-8.545 0-15.926-4.985-19.393-12.206M45.5 24H17"/><circle cx="5.783" cy="24" r="3.283" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
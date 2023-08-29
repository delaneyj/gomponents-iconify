package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleChat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="26.648" height="22.207" x="14.463" y="6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.621"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.916 13.575H9.51a2.621 2.621 0 0 0-2.622 2.621v24.753a1.048 1.048 0 0 0 1.79.742l5.91-5.91h16.328a2.621 2.621 0 0 0 2.62-2.62V16.196a2.621 2.621 0 0 0-2.62-2.621Z"/>`),
		g.Group(children),
	)
}
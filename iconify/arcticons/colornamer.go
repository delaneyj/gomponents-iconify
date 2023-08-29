package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Colornamer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/><rect width="5.5" height="7.287" x="16.305" y="20.728" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.75"/><rect width="5.5" height="7.287" x="28.453" y="20.728" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.73 23.478a2.75 2.75 0 0 1 2.75-2.75h0m-2.751 0v7.288m-12.144-11v9.625a1.375 1.375 0 0 0 1.375 1.375h.413M13.659 26.63a2.749 2.749 0 0 1-2.388 1.386h0a2.75 2.75 0 0 1-2.75-2.75v-1.788a2.75 2.75 0 0 1 2.75-2.75h0a2.749 2.749 0 0 1 2.385 1.38"/>`),
		g.Group(children),
	)
}
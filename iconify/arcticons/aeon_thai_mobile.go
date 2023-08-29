package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AeonThaiMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.274" y="5.726" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.463 23.994c-.003.622-2.947 1.125-6.577 1.125c-3.631 0-6.575-.504-6.575-1.126h0c.003-.622 2.946-1.126 6.575-1.126c3.632 0 6.576.504 6.577 1.127v0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9.5 27.669l5.642-7.338h5.542M30.451 24c0 2.214-2.092 4.008-4.673 4.008h0c-2.582 0-4.674-1.794-4.674-4.008h0c0-2.214 2.092-4.008 4.673-4.008h0c2.581 0 4.674 1.794 4.674 4.007V24Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.684 27.669h-4.347v-7.338m15.983 7.318v-7.298l6.18 7.298v-7.298"/>`),
		g.Group(children),
	)
}
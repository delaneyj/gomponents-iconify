package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Livescore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.509 21.849V9.569a2 2 0 0 1 2-2h16.347l-1.704 6.838l-8.626.013a2 2 0 0 0-2 2v2.384c-.009 1.105.895 2 2 2h7.514l2.456-.001a2 2 0 0 1 2 2l.002 8.386l.002 7.22a2 2 0 0 1-2 2l-15.724.001l1.67-6.762l8.033-.016c1.104-.002 2-.896 2-2v-2.385c.009-1.104-.895-2-2-2h0h-7.514l-2.456.002a2 2 0 0 1-2-2v-3.497m-.194 11.915l-8.39-.017a2.001 2.001 0 0 1-2-2l.013-19.808l.013-2.273c.006-1.105-.896-2-2-2l-1.403-.001h-1.05a2 2 0 0 0-1.998 2v28.84a2 2 0 0 0 2 2l13.098.024l1.717-6.766Z"/>`),
		g.Group(children),
	)
}
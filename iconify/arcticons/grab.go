package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.175 21.11a3.841 3.841 0 0 0-4.044-3.84A3.99 3.99 0 0 0 9.5 21.326v3.564a3.842 3.842 0 0 0 3.837 3.845h0a3.842 3.842 0 0 0 3.838-3.845h-3.838m6.433-.886a2.868 2.868 0 0 1 2.868-2.868h0m-2.868 2.868v4.732m10.313-2.868a2.868 2.868 0 0 1-2.868 2.868h0a2.868 2.868 0 0 1-2.868-2.868v-1.864a2.868 2.868 0 0 1 2.868-2.868h0a2.868 2.868 0 0 1 2.868 2.868m0 4.732v-7.6m2.681 2.868a2.868 2.868 0 0 1 2.868-2.868h0a2.868 2.868 0 0 1 2.868 2.868v1.864a2.868 2.868 0 0 1-2.868 2.868h0a2.868 2.868 0 0 1-2.868-2.868m0 2.868V17.264"/>`),
		g.Group(children),
	)
}
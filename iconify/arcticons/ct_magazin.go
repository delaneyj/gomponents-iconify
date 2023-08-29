package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CtMagazin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.384 21.83c.83-1.107 1.106-2.214 1.106-4.428V16.02m5.185-1.704v16.601c0 1.66 1.107 2.767 2.767 2.767h.83m-6.364-14.665h5.81M22.966 30.917c-.83 1.66-2.767 2.767-4.704 2.767h0c-3.044 0-5.534-2.49-5.534-5.533v-3.598c0-3.043 2.49-5.534 5.534-5.534h0c1.937 0 3.874 1.107 4.704 2.767"/>`),
		g.Group(children),
	)
}
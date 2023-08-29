package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Callforwardingstatus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.706 30.274a61.99 61.99 0 0 1-6.298-1.11a3.109 3.109 0 0 0-2.999.58c-.53.54-2.059 2.05-3.648 3.61a28.68 28.68 0 0 1-13.096-13.106c1.55-1.59 3-3.11 3.589-3.639a3.11 3.11 0 0 0 .58-2.999a61.812 61.812 0 0 1-1.11-6.308a2 2 0 0 0-2.299-1.769H6.908a1.5 1.5 0 0 0-1.36 1.37c-.55 7.687 3.74 15.914 4.609 17.493h0v.06l.12.23h0a35.427 35.427 0 0 0 12.995 12.996h0l.44.25h0c2 1.06 9.947 5.058 17.374 4.508a1.5 1.5 0 0 0 1.39-1.36v-8.507a2 2 0 0 0-1.77-2.299Zm-15.75-15.816l6.659.031m-.413-4.257l4.241 4.241l-4.24 4.242"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.185 9.525l4.948 4.948l-4.948 4.948"/>`),
		g.Group(children),
	)
}
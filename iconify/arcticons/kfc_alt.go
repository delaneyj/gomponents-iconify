package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KfcAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.73 25.97l-.008.05c-.388 2.201-2.487 3.986-4.688 3.986h0c-2.201 0-3.671-1.785-3.283-3.986l.716-4.06c.388-2.202 2.487-3.986 4.688-3.986h0c2.201 0 3.67 1.784 3.283 3.985l-.009.05m-24.808-4.015L10.5 30.026m.739-4.191l7.842-7.8m-2.114 11.991l-3.893-6.016m8.087-.02h3.911m-4.971 6.016l2.121-12.032h6.016"/>`),
		g.Group(children),
	)
}
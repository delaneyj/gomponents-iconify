package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m23.67 15.546l-5.431 16.806L12.5 15.546m23 6.968v5.739c0 2.357-1.947 4.202-4.202 4.202h0c-1.23 0-2.357-.41-2.971-1.23"/><path d="M35.5 15.546v6.968c0 2.357-1.947 4.202-4.202 4.202h0c-2.357 0-4.201-1.947-4.201-4.202v-6.969"/></g>`),
		g.Group(children),
	)
}
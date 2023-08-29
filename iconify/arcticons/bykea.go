package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bykea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM28.107 27.775h3.775m-3.775-7.55h3.775M28.107 24h2.461m-2.461-3.775v7.55m-5.745-7.55v7.55"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.421 27.775L23.312 24l3.109-3.75M23.312 24h-.95m-1.731-3.775L18.13 24l-2.501-3.775m2.501 7.55V24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="M12.615 24a1.888 1.888 0 1 1 0 3.775H9.5v-7.55h3.115a1.888 1.888 0 0 1 0 3.775Zm0 0H9.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.592 27.753l2.453-7.528m2.455 7.55l-2.455-7.55m1.634 5.025h-3.271"/>`),
		g.Group(children),
	)
}
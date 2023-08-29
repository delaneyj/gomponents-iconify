package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyExchange(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#4d5357" d="M25.4 15.1H8.6c-1.6 0-2.8-1.3-2.8-2.8S7 9.5 8.6 9.5h16.9c1.6 0 2.8 1.3 2.8 2.8H32c0-3.6-2.9-6.6-6.6-6.6h-6.6V2H15v3.8H8.6C4.9 5.8 2 8.7 2 12.3s2.9 6.6 6.6 6.6h16.9c1.6 0 2.8 1.3 2.8 2.8c0 1.6-1.3 2.8-2.8 2.8H8.6c-1.6 0-2.8-1.3-2.8-2.8H2c0 3.6 2.9 6.6 6.6 6.6h6.6V32H19v-3.8h6.6c3.6 0 6.6-2.9 6.6-6.6c-.2-3.5-3.1-6.5-6.8-6.5m16.7 7.6l-4.9 4.9l15.2.6l-.6-15.2l-4.9 4.9l-13-8.2zm-22.3 14l.7 15.2l4.6-4.6l13.3 7.9l-8-13.2l4.7-4.6zm33.7 6.5l8.5-8.6l-2.6-2.6l-7.7 7.9l-7.8-7.9l-2.5 2.6l8.4 8.6H46V47h3.8v3.8H46v3.7h3.8V62h3.8v-7.5h3.8v-3.7h-3.8V47h3.8v-3.8h-3.8z"/>`),
		g.Group(children),
	)
}
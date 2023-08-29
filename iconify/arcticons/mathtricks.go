package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mathtricks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.2 17.6h4m-4-6.8l2-1.1m0 0v7.9m-15.5-.7c.7.5 1.5.8 2.4.7h.3c1.4 0 2.6-1.2 2.6-2.6h0c-.1-1.4-1.2-2.5-2.6-2.6h-2.7V9.7H21m14.7 7.9h4m-4-6.8l2-1.1m0 0v7.9m-14.1-5l3 3.1m0-3l-3.1 3m.7 17.9c1.1 0 2-.9 2-2h0c0-1.1-.9-2-2-2h-1.3c-1.1 0-2 .9-2 2h0c0 1.1.9 2 2 2h0c-1.1 0-2 .9-2 2h0c0 1.1.9 2 2 2h1.3c1.1 0 2-.9 2-2h0c0-1.1-.9-2-2-2m4.3 4.4c.7.6 1.5.8 2.4.7h.3c1.4 0 2.6-1.2 2.6-2.6h0c0-1.4-1.2-2.6-2.6-2.6h-2.7v-2.8h5.3"/><circle cx="36.7" cy="38.7" r=".8" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.7 29.8v6.4m3.8 6.3h-33c-1.1 0-2-.9-2-2v-33c0-1.1.9-2 2-2h33c1.1 0 2 .9 2 2v33c0 1.1-.9 2-2 2zM20.8 22.1h6.4m-6.4 3.8h6.4m-4.3 7.7h1.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.6 13.7c1.1 0 2-.9 2-2h0c0-1.1-.9-2-2-2c0 0 0 0 0 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.6 17.6c1.1 0 2-.9 2-2h0c0-1.1-.9-2-2-2c0 0 0 0 0 0M8.3 17c.7.5 1.6.8 2.5.7h.8m-3.3-7.3c.7-.6 1.6-.8 2.5-.7h.8m-2 4h2m4.9 20.9c1.1 0 2-.9 2-2h0c0-1.1-.9-2-2-2c0 0 0 0 0 0m0 8c1.1 0 2-.9 2-2c0 0 0 0 0 0h0c0-1.1-.9-2-2-2c0 0 0 0 0 0m-3.3 3.3c.7.5 1.6.8 2.5.7h.8m-3.3-7.3c.7-.5 1.6-.8 2.5-.7h.8m-2 4h2"/>`),
		g.Group(children),
	)
}
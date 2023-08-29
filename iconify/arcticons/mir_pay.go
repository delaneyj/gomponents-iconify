package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirPay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM30.436 33.873l-2.1-5.3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.336 28.573l-2.5 7.2c-.2.5-.6.8-1.1.8h-.3m-2.036-4.75c0 1.1-.9 2-2 2h0c-1.1 0-2-.9-2-2v-1.3c0-1.1.9-2 2-2h0c1.1 0 2 .9 2 2m0 3.3v-5.3m-10.736 5.3v-8h2.6c1.5 0 2.7 1.2 2.7 2.7s-1.2 2.7-2.7 2.7h-2.6m14.192-7.851V11.427h3.882c2.24 0 4.032 1.791 4.032 4.031s-1.792 4.032-4.032 4.032h-3.882m4.126-.005l3.788 3.887M26.016 11.427v11.945m-15.786 0V11.427l5.973 11.945l5.972-11.945v11.945"/>`),
		g.Group(children),
	)
}
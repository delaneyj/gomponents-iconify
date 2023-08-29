package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessPgnMasterPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.6 14.4v3.5m0-3.5v-1.7h.7H14h.6m1.6 1.7H13"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2ZM24 5.5v37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.2 38.2c.2 0 .3-.1.3-.3h0c0-1.9-2.5-2-1.2-2.2c.7-.1 1.7-.8.8-1.7s-1.2-.7-.9-1s2-2 .9-2.3c-1.6-.5-1-.7-.9-.8c.9-.4 1.5-1.3 1.6-2.3c0-.1-.1-.3-.2-.3s-.2 0-.2.1c-.4.5-4 1.7-9.4 1.7c-.8.1-.4.9-.3 1.8c.1.5.9.6 1.3.6c.2 0 1.2-.1 1.4-.1c.6.1.8.6 1 1c.4.6 1.2 1 1.9.9c.1 0-.2.2-1.1.1c-1.5 0-3.5 1.7-4.1 4.2c0 .2.1.4.3.5h0c1.5.1 7.2 0 8.8.1Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.1 30.4c.3 0 .6 0 .9-.1c-.3 0-.6 0-.9.1Zm-18.5-10h-4.3s-3.9-7 4.2-11.8h0c8.1 4.8 4.3 11.8 4.3 11.8h-4.2ZM5.5 24h37m-5-11.2c.1 1.5-1 2.9-2.5 3c-.8.1-1.6-.2-2.1-.7c-.5-.4-.8-1.1-.9-1.8c-.1-1.5 1-2.9 2.5-3s2.8.9 3 2.5Zm-4.9 2.3l-3 3.3l-.9 1m2.4.3l-1.5-1.3"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsnBankieren(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.4 43.5c-7.1-.1-12.8-5.7-13.1-12.7c0-1.1.1-2.2.3-3.3c.2-1.2.7-2.3 1.2-3.4c.3-.6.6-1.3.7-2c.1-.3.1-.6.1-.9c0-2.8-2.2-5.2-5.1-5.2c-2.1 0-4 1.2-4.8 3.2c-.2-.9-.4-1.9-.3-2.9c-.3-6.4 4.7-11.7 11-12c6.4-.3 11.7 4.7 12 11v.5c.1 2.6-.7 5.2-2.2 7.3c-1.8 2.8-3.9 5.4-3.9 9c0 5 3.2 9.5 8 11c-1.3.3-2.6.4-3.9.4h0zm13.4-13.1c-2.3 1-3.4 3.8-2.3 6.1c.1.3.3.5.4.7c-1.4 2.2-3.4 4-5.7 5.1c-2.5-3.4-1.7-8.2 1.8-10.7c1.3-.9 2.9-1.4 4.5-1.4c.4 0 .9 0 1.3.2h0zm.7-7.8c-.2 1.2-.6 2.3-1.3 3.2c-.7-2-1.8-3.8-3.3-5.2c-.4-.4-.8-.7-1.2-1h0c-2.1-1.5-3.3-3.9-3.3-6.5c0-.8.1-1.6.3-2.3c.4 1.3 1.2 2.5 2.4 3.2l.9.6l1.4.6h0c1-1.2 1.6-2.8 1.8-4.3c1.1 1.8 1.2 3.9.5 5.9h0c1.3 1.2 2.1 2.9 1.9 4.7c0 .4 0 .8-.1 1.1h0z"/><circle cx="33.7" cy="21.5" r=".8" fill="currentColor"/>`),
		g.Group(children),
	)
}
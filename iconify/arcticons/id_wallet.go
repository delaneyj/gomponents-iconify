package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="36.65" height="26.043" x="5.675" y="10.979" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3"/><circle cx="14.838" cy="21.487" r="3.563" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.328 31.095h9.783a.92.92 0 0 0 .7-1.52a7.172 7.172 0 0 0-11.183 0a.92.92 0 0 0 .7 1.52ZM28.709 20.85h6.999m-6.999 6.872h6.999m-6.999-3.436h9.671"/>`),
		g.Group(children),
	)
}
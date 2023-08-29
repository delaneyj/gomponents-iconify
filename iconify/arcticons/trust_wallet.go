package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrustWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c1.35-.062 8.977-4.15 11.767-8.742c2.79-4.592 4.268-16.815 4.37-24.879C32.072 9.88 26.853 7.352 24 4.5c-2.855 2.855-8.062 5.38-16.138 5.38c0 8.065 1.68 20.217 4.371 24.878C14.924 39.42 22.65 43.561 24 43.5Z"/>`),
		g.Group(children),
	)
}
package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dkbbanking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.7 29.54V18.42m0 0h4.68a2.79 2.79 0 0 1 2.79 2.79h0A2.79 2.79 0 0 1 34.38 24H29.7m0 0h4.68a2.79 2.79 0 0 1 2.79 2.79h0a2.79 2.79 0 0 1-2.79 2.79H29.7m-18.87 0V18.42h1.89A5.58 5.58 0 0 1 18.3 24h0a5.58 5.58 0 0 1-5.58 5.58Zm10.39-11.16v11.15m1.44-5.58l4.12-5.54m-4.12 5.54l4.12 5.59m-4.12-5.59h-1.44"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}
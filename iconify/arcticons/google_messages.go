package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleMessages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.8 29.286c-6.247 0-11.3-5.33-11.3-11.893S10.553 5.5 16.8 5.5h11.757c6.247 0 11.3 5.33 11.3 11.893s-5.053 11.893-11.3 11.893H16.8v0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.42 34.092v7.252c0 .63-.508 1.156-1.117 1.156a1.096 1.096 0 0 1-.788-.342l-11.02-11.35c-2.159-2.208-3.352-5.23-3.352-8.356c0-6.438 5.053-11.667 11.3-11.667H31.2c6.247 0 11.3 5.23 11.3 11.667s-5.053 11.666-11.3 11.666h-6.78v-.026Z"/>`),
		g.Group(children),
	)
}
package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForAruba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#42ade2" d="M59.5 44v-4h1.4c.7-2.5 1.1-5.2 1.1-8C62 15.4 48.6 2 32 2S2 15.4 2 32c0 2.8.4 5.5 1.1 8h1.4v4c.9 2.1 2.1 4.2 3.5 6h3.6v4c5.4 5 12.5 8 20.4 8c7.9 0 15-3 20.4-8v-4H56c1.4-1.8 2.6-3.9 3.5-6"/><path fill="#ffe62e" d="M4.5 44h55c.6-1.3 1-2.6 1.4-4H3.1c.4 1.4.8 2.7 1.4 4M8 50c1.1 1.4 2.3 2.8 3.6 4h40.8c1.3-1.2 2.5-2.6 3.6-4H8z"/><path fill="#fff" d="m22 33.7l-3.7-10l-10-3.7l10-3.7l3.7-10l3.7 10l10 3.7l-10 3.7l-3.7 10"/><path fill="#c94747" d="M19.8 17.8L22 12l2.2 5.8L30 20l-5.8 2.2L22 28l-2.2-5.8L14 20z"/>`),
		g.Group(children),
	)
}
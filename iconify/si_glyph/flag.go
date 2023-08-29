package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 0h1.994v15.913H1zm3.056.52v7.575S5.667 6.664 9.244 8c3.576 1.338 4.305.974 5.712.742c0 0-2.048-.871-3.222-4.029c0 0 2.987-2.755 3.222-4.274c0 0-3.7 1.212-5.751.241C7.152-.293 4.994-.089 4.056.52z"/>`),
		g.Group(children),
	)
}
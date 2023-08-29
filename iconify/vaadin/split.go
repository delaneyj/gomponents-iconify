package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Split(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 11h6v5H0v-5zm11-1V8l-.64.64a4.427 4.427 0 0 1-1.358-3.658L11.001 5V0h-6v5h2a4.43 4.43 0 0 1-1.358 3.639l-.642-.638v2h2l-.65-.65A5.426 5.426 0 0 0 8 4.982c-.01.149-.016.299-.016.45c0 1.539.639 2.928 1.665 3.917l-.648.652h2zm-1 1h6v5h-6v-5z"/>`),
		g.Group(children),
	)
}
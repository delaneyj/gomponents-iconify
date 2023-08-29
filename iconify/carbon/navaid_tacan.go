package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavaidTacan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="14" r="2" fill="currentColor"/><path fill="currentColor" d="M20 28h-8a1 1 0 0 1-1-1v-5.734l-3.739-6.543l-4.776-2.866a1 1 0 0 1-.353-1.353l4-7a1 1 0 0 1 1.383-.361L12.277 6h7.446l4.762-2.857a1 1 0 0 1 1.383.36l4 7a1 1 0 0 1-.353 1.354l-4.776 2.866L21 21.266V27a1 1 0 0 1-1 1Zm-7-2h6v-5.265l4.262-7.458l4.384-2.63l-3.007-5.264L20.277 8h-8.554L7.361 5.383l-3.007 5.263l4.384 2.63L13 20.736Z"/>`),
		g.Group(children),
	)
}
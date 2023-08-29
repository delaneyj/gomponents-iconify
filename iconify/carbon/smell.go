package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 15v-5h-2v7h2a3 3 0 0 1 0 6h-1v-2h-2v2a3 3 0 0 1-6 0v-2h-2v2h-1a3 3 0 0 1 0-6h2V9a3 3 0 0 1 3-3h1V4h-1a5 5 0 0 0-5 5v6a5 5 0 0 0 0 10h1.42a5 5 0 0 0 9.16 0H22a5 5 0 0 0 0-10Z"/>`),
		g.Group(children),
	)
}
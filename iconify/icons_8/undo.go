package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m12.78 5.28l-8 8l-.686.72l.687.72l8 8l1.44-1.44L7.936 15H21c2.755 0 5 2.245 5 5v7h2v-7c0-3.845-3.155-7-7-7H7.937l6.282-6.28l-1.44-1.44z"/>`),
		g.Group(children),
	)
}
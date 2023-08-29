package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.095 8.43l-1.424-1.404l-4.914 4.985l4.985 4.914l1.404-1.424l-2.502-2.467l6.497.05l.016-2l-6.628-.05l2.566-2.604ZM5.467 15.562l1.416 1.412l4.944-4.956l-4.956-4.943L5.459 8.49l2.591 2.585l-7.206.024l.006 2l7.097-.024l-2.48 2.486Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 2C7.373 2 2 7.373 2 14s5.373 12 12 12s12-5.373 12-12S20.627 2 14 2Zm-1.234 7.278l6.505 3.862a1 1 0 0 1 0 1.72l-6.505 3.862a1.5 1.5 0 0 1-2.266-1.29v-6.864a1.5 1.5 0 0 1 2.266-1.29Z"/>`),
		g.Group(children),
	)
}
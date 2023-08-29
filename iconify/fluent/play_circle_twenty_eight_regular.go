package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.766 9.278a1.5 1.5 0 0 0-2.266 1.29v6.864a1.5 1.5 0 0 0 2.266 1.29l6.505-3.862a1 1 0 0 0 0-1.72l-6.505-3.862ZM2 14C2 7.373 7.373 2 14 2s12 5.373 12 12s-5.373 12-12 12S2 20.627 2 14ZM14 3.5C8.201 3.5 3.5 8.201 3.5 14S8.201 24.5 14 24.5S24.5 19.799 24.5 14S19.799 3.5 14 3.5Z"/>`),
		g.Group(children),
	)
}
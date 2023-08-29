package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOutlineDownLeftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.151 1.44a1.5 1.5 0 0 0-2.121 0L5.004 5.464L3.907 4.368c-.75-.75-2.033-.285-2.13.77l-.772 8.5a1.25 1.25 0 0 0 1.358 1.358l8.499-.773c1.055-.096 1.52-1.379.77-2.128l-1.098-1.099L14.56 6.97a1.5 1.5 0 0 0 0-2.121l-3.409-3.41Z"/>`),
		g.Group(children),
	)
}
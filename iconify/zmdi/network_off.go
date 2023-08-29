package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M427 5v367L243 189zM80 80l368 368l-27 27l-43-43H0l189-189L53 107z"/>`),
		g.Group(children),
	)
}
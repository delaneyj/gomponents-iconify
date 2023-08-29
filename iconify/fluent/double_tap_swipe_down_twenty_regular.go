package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleTapSwipeDownTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.354 17.854a.5.5 0 0 1-.708 0l-3-3a.5.5 0 0 1 .708-.708L9.5 16.293V7.5a.5.5 0 0 1 1 0v8.793l2.146-2.147a.5.5 0 0 1 .708.708l-3 3ZM4.5 7.5a5.502 5.502 0 0 0 4 5.293v-1.05A4.502 4.502 0 0 1 10 3a4.5 4.5 0 0 1 1.5 8.744v1.049a5.5 5.5 0 1 0-7-5.293Zm2 0a3.5 3.5 0 0 0 2 3.163V9.5a2.5 2.5 0 1 1 3 0v1.163a3.5 3.5 0 1 0-5-3.163Z"/>`),
		g.Group(children),
	)
}
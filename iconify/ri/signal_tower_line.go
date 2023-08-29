package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalTowerLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.116 20.087l1.015-1.739a8 8 0 1 1 9.738 0l1.015 1.739A9.986 9.986 0 0 0 22 12c0-5.523-4.477-10-10-10S2 6.477 2 12a9.986 9.986 0 0 0 4.116 8.087Zm2.034-3.485a6 6 0 1 1 7.7 0l-1.03-1.766a4 4 0 1 0-5.64 0l-1.03 1.766ZM11 13h2v9h-2v-9Z"/>`),
		g.Group(children),
	)
}
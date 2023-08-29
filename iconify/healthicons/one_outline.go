package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24.879 9.879A3 3 0 0 1 30 12v24a3 3 0 0 1-6 0V19.236a3 3 0 0 1-4.121-4.357l5-5Zm2.504 1.197a1 1 0 0 0-1.09.217l-5 5a1 1 0 0 0 1.414 1.414l1.586-1.586a1 1 0 0 1 1.707.707V36a1 1 0 1 0 2 0V12a1 1 0 0 0-.617-.924Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
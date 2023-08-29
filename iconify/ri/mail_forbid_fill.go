package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailForbidFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.266 11.554l4.388-3.798l-1.308-1.512l-6.285 5.439l-6.414-5.445l-1.294 1.524l7.702 6.54A6.967 6.967 0 0 0 11 18c0 1.074.242 2.09.674 3H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v8.255A6.968 6.968 0 0 0 18 11a6.98 6.98 0 0 0-2.734.554Zm1.44 9.154a3 3 0 0 0 4.001-4.001l-4 4Zm-1.414-1.415l4.001-4a3 3 0 0 0-4.001 4.001ZM18 23a5 5 0 1 1 0-10a5 5 0 0 1 0 10Z"/>`),
		g.Group(children),
	)
}
package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailForbidLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 7.238l-7.928 7.1L4 7.216V19h7.07c.102.706.308 1.378.604 2H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v8.255a6.972 6.972 0 0 0-2-.965V7.238ZM19.501 5H4.511l7.55 6.662L19.502 5Zm-2.794 15.708a3 3 0 0 0 4.001-4.001l-4.001 4Zm-1.415-1.415l4.001-4a3 3 0 0 0-4.001 4.001ZM18 23a5 5 0 1 1 0-10a5 5 0 0 1 0 10Z"/>`),
		g.Group(children),
	)
}
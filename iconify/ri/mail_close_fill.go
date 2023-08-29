package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailCloseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 13.341A6 6 0 0 0 14.341 21H3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v9.341Zm-9.94-1.658L5.648 6.238L4.353 7.762l7.72 6.555l7.581-6.56l-1.308-1.513l-6.285 5.439ZM21.415 19l2.121 2.121l-1.414 1.415L20 20.413l-2.121 2.121l-1.415-1.414L18.587 19l-2.121-2.121l1.414-1.415L20 17.587l2.121-2.121l1.415 1.414L21.413 19Z"/>`),
		g.Group(children),
	)
}
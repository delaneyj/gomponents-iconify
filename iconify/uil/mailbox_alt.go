package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailboxAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13h2a1 1 0 0 0 0-2h-2a1 1 0 0 0 0 2Zm8 7h-1V9h1a1 1 0 0 0 0-2h-1.09A6 6 0 0 0 6.09 7H5a1 1 0 0 0 0 2h1v11H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2ZM12 4a4 4 0 0 1 3.86 3H8.14A4 4 0 0 1 12 4Zm4 16H8v-2h8Zm0-4H8V9h8Z"/>`),
		g.Group(children),
	)
}
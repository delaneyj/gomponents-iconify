package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailboxUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 9h4v2H5V9m17 0v9a2 2 0 0 1-2 2H2V9a5 5 0 0 1 5-5h10a5 5 0 0 1 5 5M10 9a3 3 0 0 0-3-3a3 3 0 0 0-3 3v9h6V9m6-2h-4v6h2V9h2V7Z"/>`),
		g.Group(children),
	)
}
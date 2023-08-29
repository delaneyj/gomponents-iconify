package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mailbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 12h2a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2Zm9-6h-5V4h1a1 1 0 0 0 0-2h-2a1 1 0 0 0-1 1v3H7a4 4 0 0 0-4 4v6a1 1 0 0 0 1 1h6v4a1 1 0 0 0 2 0v-4h8a1 1 0 0 0 1-1v-6a4 4 0 0 0-4-4Zm-4 4v5H5v-5a2 2 0 0 1 2-2h6.56a3.91 3.91 0 0 0-.56 2Zm6 5h-4v-5a2 2 0 0 1 4 0Z"/>`),
		g.Group(children),
	)
}
package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrapText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2Zm6 8H3a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Zm9.5-5H3a1 1 0 0 0 0 2h15.5a1.5 1.5 0 0 1 0 3h-3.09l.3-.29a1 1 0 0 0-1.42-1.42l-2 2a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l2 2a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.3-.29h3.09a3.5 3.5 0 0 0 0-7Z"/>`),
		g.Group(children),
	)
}
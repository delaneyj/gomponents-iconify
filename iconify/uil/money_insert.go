package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyInsert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.46 6l.54-.59V9a1 1 0 0 0 2 0V5.41l.54.55A1 1 0 0 0 15 6a1 1 0 0 0 0-1.42l-2.29-2.29a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21L9 4.54A1 1 0 0 0 10.46 6ZM12 12a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm-7-1a1 1 0 1 0 1-1a1 1 0 0 0-1 1Zm14 0a1 1 0 1 0-1 1a1 1 0 0 0 1-1Zm1-7h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h4a1 1 0 0 0 0-2H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}
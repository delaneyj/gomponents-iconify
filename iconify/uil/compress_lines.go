package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompressLines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20h-4v-3.59l.79.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-2.5-2.5a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2.5 2.5a1 1 0 0 0 1.42 1.42l.79-.8V20H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2ZM7 4h4v3.59l-.79-.8a1 1 0 1 0-1.42 1.42l2.5 2.5a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2.5-2.5a1 1 0 1 0-1.42-1.42l-.79.8V4h4a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 0l-.83 5H7a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h.17L8 24h8l.83-5H17a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-.17L16 0H8M7 7h10v10H7V7Z"/>`),
		g.Group(children),
	)
}
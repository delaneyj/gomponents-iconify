package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatioHeater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 22H9v-1h6v1m4-18l-4-2H9L5 4h14M8 5l.4 1h7.2l.4-1H8m2 5h1v5c-.6 0-1 .4-1 1v4h4v-4c0-.6-.4-1-1-1v-5h1l.4-1H9.6l.4 1m-.8-2h5.6l.4-1H8.8l.4 1Z"/>`),
		g.Group(children),
	)
}
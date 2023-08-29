package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowOpenVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 20V2H3v18H1v3h22v-3M19 4v7h-2V4M5 4h2v7H5m0 9v-7h2v7m2 0V4h6v16m2 0v-7h2v7Z"/>`),
		g.Group(children),
	)
}
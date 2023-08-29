package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3c.53 0 1.04.21 1.41.59c.38.37.59.88.59 1.41v10c0 .53-.21 1.04-.59 1.41c-.37.38-.88.59-1.41.59H7c-.53 0-1.04-.21-1.41-.59C5.21 16.04 5 15.53 5 15V5c0-.53.21-1.04.59-1.41C5.96 3.21 6.47 3 7 3h14M3 19h15v2H3c-.53 0-1.04-.21-1.41-.59C1.21 20.04 1 19.53 1 19V8h2v11Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckbookArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 14h10.53c-.28.31-.53.64-.73 1H5v-1m16-6v4.08c.72.12 1.39.37 2 .72V5H1v14h13.08a6.53 6.53 0 0 1-.08-1c0-.34.03-.67.08-1H3V8h18M5 10h7v2H5v-2m11 8l3-3v2h4v2h-4v2l-3-3Z"/>`),
		g.Group(children),
	)
}
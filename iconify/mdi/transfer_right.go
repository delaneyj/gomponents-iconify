package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransferRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 8h2v8H3V8m4 0h2v8H7V8m4 0h2v8h-2V8m4 11.25V4.75L22.25 12L15 19.25Z"/>`),
		g.Group(children),
	)
}
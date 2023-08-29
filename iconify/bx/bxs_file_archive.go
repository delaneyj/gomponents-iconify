package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsFileArchive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M6 2a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8l-6-6h-3v2H9v2h2v2H9v2h2v8H7v-6h2v-2H7V8h2V6H7V4h2V2H6zm7 2l5 5h-5V4z" fill="currentColor"/><path d="M8 15h2v2H8z" fill="currentColor"/>`),
		g.Group(children),
	)
}
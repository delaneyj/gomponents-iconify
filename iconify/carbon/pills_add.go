package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 14a7.94 7.94 0 0 0-4 1.083V9A7 7 0 0 0 4 9v14a6.986 6.986 0 0 0 12.276 4.577A7.997 7.997 0 1 0 22 14ZM11 4a5.006 5.006 0 0 1 5 5v6H6V9a5.006 5.006 0 0 1 5-5Zm0 24a5.006 5.006 0 0 1-5-5v-6h9.765a7.956 7.956 0 0 0-.724 8.932A5.015 5.015 0 0 1 11 28Zm11 0a6 6 0 1 1 6-6a6.007 6.007 0 0 1-6 6Z"/><path fill="currentColor" d="M25 21h-2v-2h-2v2h-2v2h2v2h2v-2h2v-2z"/>`),
		g.Group(children),
	)
}
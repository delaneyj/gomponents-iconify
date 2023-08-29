package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m23 27.18l-2.59-2.59L19 26l4 4l7-7l-1.41-1.41L23 27.18z"/><circle cx="11" cy="8" r="1" fill="currentColor"/><circle cx="11" cy="16" r="1" fill="currentColor"/><circle cx="11" cy="24" r="1" fill="currentColor"/><path fill="currentColor" d="M24 3H8a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h8v-2H8v-6h18V5a2 2 0 0 0-2-2Zm0 16H8v-6h16Zm0-8H8V5h16Z"/>`),
		g.Group(children),
	)
}
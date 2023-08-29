package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 27H8c-1.5 0-3-1.5-3-3v-8m0 0h5.5s1 3.5 5.5 3.5s5.5-3.5 5.5-3.5H27L23 5H9L5 16Zm20 12v-8m-4 4h8"/>`),
		g.Group(children),
	)
}
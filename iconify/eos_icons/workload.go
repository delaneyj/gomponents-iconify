package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Workload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6H2v14a2 2 0 0 0 2 2h14v-2H4Z"/><path fill="currentColor" d="M10 10h8v2h-8zm0 3h8v2h-8zm3.47-8.01h-2.31L10 7l1.16 2h2.31l1.16-2l-1.16-2.01z"/><path fill="currentColor" d="M20 2H8a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm0 14H8V4h12Z"/>`),
		g.Group(children),
	)
}
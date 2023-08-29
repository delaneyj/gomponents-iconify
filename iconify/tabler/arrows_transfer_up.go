package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsTransferUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21v-6m13-9l-3-3l-3 3m3-3v18m-7-3l-3 3l-3-3M7 3v2m0 4v2"/>`),
		g.Group(children),
	)
}
package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h2v2H4V4Zm0 14h2v2H4v-2ZM18 4h2v2h-2V4Zm0 7h2v2h-2v-2Zm-7 0h2v2h-2v-2Zm-7 0h2v2H4v-2Zm7-7h2v2h-2V4Zm0 14h2v2h-2v-2Zm7 0h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HttpDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8v8h2a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2H3zm11 0h-4v8h4m-4-4h2.5M17 8v8h4"/>`),
		g.Group(children),
	)
}
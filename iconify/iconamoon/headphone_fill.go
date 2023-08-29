package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphoneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 4a8 8 0 0 0-8 8v2h2a2 2 0 0 1 2 2v3a3 3 0 1 1-6 0v-7C2 6.477 6.477 2 12 2s10 4.477 10 10v7a3 3 0 1 1-6 0v-3a2 2 0 0 1 2-2h2v-2a8 8 0 0 0-8-8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
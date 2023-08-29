package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 2a1 1 0 1 0-2 0v1H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h5v1a1 1 0 0 0 2 0v-1a1 1 0 0 0 0-2V5a1 1 0 1 0 0-2V2Zm4 1a1 1 0 1 0 0 2h1a1 1 0 0 1 1 1v1a1 1 0 1 0 2 0V6a3 3 0 0 0-3-3h-1Zm4 8a1 1 0 1 0-2 0v2a1 1 0 0 0 2 0v-2Zm0 6a1 1 0 1 0-2 0v1a1 1 0 0 1-1 1h-1a1 1 0 1 0 0 2h1a3 3 0 0 0 3-3v-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
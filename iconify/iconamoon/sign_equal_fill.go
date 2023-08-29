package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignEqualFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 8a1 1 0 0 0 0 2h12a1 1 0 1 0 0-2H6Zm0 6a1 1 0 1 0 0 2h12a1 1 0 1 0 0-2H6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
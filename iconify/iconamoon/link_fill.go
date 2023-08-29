package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 9a3 3 0 1 0 0 6h3a1 1 0 1 1 0 2H6A5 5 0 0 1 6 7h3a1 1 0 0 1 0 2H6Zm1 3a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1Zm8-5a1 1 0 1 0 0 2h3a3 3 0 1 1 0 6h-3a1 1 0 1 0 0 2h3a5 5 0 0 0 0-10h-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.947 5a1.5 1.5 0 0 1 1.5-1.5h13a1.5 1.5 0 0 1 1.5 1.5v4a1.5 1.5 0 0 1-1.5 1.5h-13a1.5 1.5 0 0 1-1.5-1.5V5ZM6 7a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm4-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-6.053 7a1.5 1.5 0 0 1 1.5-1.5h13a1.5 1.5 0 0 1 1.5 1.5v4a1.5 1.5 0 0 1-1.5 1.5H12.75v1.75H19a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1 0-1.5h6.25V18.5H5.447a1.5 1.5 0 0 1-1.5-1.5v-4ZM6 15a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm3 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
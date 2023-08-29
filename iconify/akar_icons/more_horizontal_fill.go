package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreHorizontalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm8 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm8 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
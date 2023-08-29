package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVerticalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm0 8a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm0 8a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
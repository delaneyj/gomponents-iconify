package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3-17a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm0 8a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-3 11a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IfDesign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 5h4v14h-4V5ZM5 19v-9h4v9H5ZM7 5a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm8 0h4v4h-4V5Zm4 5h-4v4h4v-4Z"/>`),
		g.Group(children),
	)
}
package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PointerUpFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20.05 6.35a3 3 0 0 1 1.2 2.4v11a3 3 0 0 1-3 3h-12a3 3 0 0 1-3-3v-11a3 3 0 0 1 1.2-2.4l6-4.5a3 3 0 0 1 3.6 0l6 4.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
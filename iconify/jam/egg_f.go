package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EggF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20.565a8 8 0 0 1-8-8c0-4.418 3.582-12 8-12s8 7.582 8 12a8 8 0 0 1-8 8z"/>`),
		g.Group(children),
	)
}
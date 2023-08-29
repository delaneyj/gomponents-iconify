package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Merck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 6a6 6 0 0 1 12 0zm0 12a6 6 0 0 1 6-6a6 6 0 0 1-6-6a6 6 0 0 0 0 12a6 6 0 1 0 12 0zm6-6a6 6 0 0 1 6 6a6 6 0 1 0 0-12a6 6 0 0 1-6 6"/>`),
		g.Group(children),
	)
}
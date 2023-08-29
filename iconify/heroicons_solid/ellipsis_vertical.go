package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 3a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm0 5.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm1.5 7a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Z"/>`),
		g.Group(children),
	)
}
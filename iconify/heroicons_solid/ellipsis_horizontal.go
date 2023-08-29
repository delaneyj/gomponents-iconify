package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 10a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm5.5 0a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Zm7-1.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/>`),
		g.Group(children),
	)
}
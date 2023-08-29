package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OtherTwoSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-width="1.5" d="M12 5.25a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Zm0 6a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Zm0 6a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5Z"/>`),
		g.Group(children),
	)
}
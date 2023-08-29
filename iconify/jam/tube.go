package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7a1 1 0 1 1 0 2v2a1 1 0 0 1 0 2v4a1 1 0 0 0 2 0V2H2v5zM0 0h6v17a3 3 0 0 1-6 0V0z"/>`),
		g.Group(children),
	)
}
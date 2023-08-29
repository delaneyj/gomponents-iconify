package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransformTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 1v15h6v2H6V8H1V6h5V1h2Zm2 5h8v10h5v2h-5v5h-2V8h-6V6Z"/>`),
		g.Group(children),
	)
}
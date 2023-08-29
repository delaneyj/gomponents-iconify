package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 16.75H6a.76.76 0 0 1-.67-.41a.75.75 0 0 1 .07-.79l6-8a.77.77 0 0 1 1.2 0l6 8a.75.75 0 0 1 .07.79a.76.76 0 0 1-.67.41Zm-10.5-1.5h9l-4.5-6Z"/>`),
		g.Group(children),
	)
}
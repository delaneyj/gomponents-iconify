package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16.75a.74.74 0 0 1-.6-.3l-6-8a.75.75 0 0 1-.07-.79a.76.76 0 0 1 .67-.41h12a.76.76 0 0 1 .67.41a.75.75 0 0 1-.07.79l-6 8a.74.74 0 0 1-.6.3Zm-4.5-8l4.5 6l4.5-6Z"/>`),
		g.Group(children),
	)
}
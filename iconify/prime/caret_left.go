package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 18.75a.74.74 0 0 1-.45-.15l-8-6a.75.75 0 0 1 0-1.2l8-6a.75.75 0 0 1 .79-.07a.76.76 0 0 1 .41.67v12a.76.76 0 0 1-.41.67a.84.84 0 0 1-.34.08ZM9.25 12l6 4.5v-9Z"/>`),
		g.Group(children),
	)
}
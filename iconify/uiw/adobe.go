package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20 1v18.001L12.607 1H20ZM7.399 1L0 19.001V1h7.399Zm2.604 6.265L14.713 19h-3.086l-1.41-3.419H6.77l3.233-8.316Z"/>`),
		g.Group(children),
	)
}
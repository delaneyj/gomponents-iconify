package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 6.086l9.414 9.414H2.586L12 6.086ZM7.414 13.5h9.172L12 8.914L7.414 13.5Z"/>`),
		g.Group(children),
	)
}
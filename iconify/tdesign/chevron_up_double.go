package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUpDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.5 18.414l-4.5-4.5l-4.5 4.5L6.086 17L12 11.086L17.914 17L16.5 18.414Zm0-6.5l-4.5-4.5l-4.5 4.5L6.086 10.5L12 4.586l5.914 5.914l-1.414 1.414Z"/>`),
		g.Group(children),
	)
}
package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDownDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.5 5.586l-4.5 4.5l-4.5-4.5L6.086 7L12 12.914L17.914 7L16.5 5.586Zm0 6.5l-4.5 4.5l-4.5-4.5L6.086 13.5L12 19.414l5.914-5.914l-1.414-1.414Z"/>`),
		g.Group(children),
	)
}
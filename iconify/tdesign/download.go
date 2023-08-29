package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 3v9.586l3.5-3.5l1.414 1.414L12 16.414L6.086 10.5L7.5 9.086l3.5 3.5V3h2ZM4.5 14v5h15v-5h2v7h-19v-7h2Z"/>`),
		g.Group(children),
	)
}
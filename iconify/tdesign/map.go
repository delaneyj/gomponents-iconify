package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5v15.574l-7 4.084l-6.074-3.544L2 21.5V5.926l7-4.084Zm1 15.084l4 2.333V7.074l-4-2.333v12.185ZM8 4.74L4 7.074V18.5l4-1.667V4.741Zm8 2.426v12.092l4-2.333V5.5l-4 1.667Z"/>`),
		g.Group(children),
	)
}
package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.886 1.859l8.086 3.537L22 2.382V11h-2V5.618l-3 1.5V11h-2V7.154L9 4.53v10.853l2.34 1.17l-.895 1.789l-2.401-1.201L2 20.766V5.98l5.886-4.12ZM7 15.434V4.92l-3 2.1v10.213l3-1.8ZM19 12v4h4v2h-4v4h-2v-4h-4v-2h4v-4h2Z"/>`),
		g.Group(children),
	)
}
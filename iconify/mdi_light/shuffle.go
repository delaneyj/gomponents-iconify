package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shuffle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 5h6v6h-1V6.71L4.71 20L4 19.29L17.29 6H13V5m0 14h4.29l-4.58-4.59l.7-.7L18 18.29V14h1v6h-6v-1M4 5.71L4.71 5l5.58 5.59l-.7.7L4 5.71Z"/>`),
		g.Group(children),
	)
}
package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ingress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 24h-4v-4h4v4Zm2-13.982h-3v7.996h-2v-7.996H8l1-1.001l1-1.001l1-1.001l1-1.001l1 1l1 1.002l1 1ZM1 15.999V3.001a1.983 1.983 0 0 1 .158-.777a2.02 2.02 0 0 1 1.065-1.066A1.981 1.981 0 0 1 3 1h18a1.981 1.981 0 0 1 .777.158a2.02 2.02 0 0 1 1.065 1.066A1.983 1.983 0 0 1 23 3v12.998a1.983 1.983 0 0 1-.158.777a2.02 2.02 0 0 1-1.065 1.066A1.981 1.981 0 0 1 21 18h-6v-1.99h6v-13H3v13h6V18H3a1.981 1.981 0 0 1-.777-.158a2.02 2.02 0 0 1-1.065-1.066A1.983 1.983 0 0 1 1 16Z"/>`),
		g.Group(children),
	)
}
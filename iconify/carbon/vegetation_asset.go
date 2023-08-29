package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VegetationAsset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="21.5" cy="10.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="M28.5 1a2.452 2.452 0 0 0-1.206.31l-8.31 4.298l.002.006a5.497 5.497 0 1 0 7.326 7.544l.003.002l4.4-8.5A2.5 2.5 0 0 0 28.5 1zm-7 13a3.5 3.5 0 1 1 3.5-3.5a3.504 3.504 0 0 1-3.5 3.5zm7.441-10.265l-2.384 4.606a5.528 5.528 0 0 0-2.899-2.899l4.595-2.377A.516.516 0 0 1 29 3.5a.498.498 0 0 1-.059.235zM19 20h7v-2h-9v10a2.003 2.003 0 0 0 2 2h3v-2h-3zm-7 10H9v-2h3V15.566l-3.515-2.109l1.03-1.714l3.514 2.108A2.011 2.011 0 0 1 14 15.566V28a2.002 2.002 0 0 1-2 2z"/><path fill="currentColor" d="M17.308 2.285A9.488 9.488 0 0 0 15 2a8.027 8.027 0 0 0-6.921 4.001L8 6a6 6 0 0 0 0 12v-2a4 4 0 0 1 0-8a2.7 2.7 0 0 1 .387.04l.863.113l.309-.667A6.02 6.02 0 0 1 15 4a7.499 7.499 0 0 1 1.731.215Z"/>`),
		g.Group(children),
	)
}
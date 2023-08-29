package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SawtoothWave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22V6.83L2 16v-2.83L13 2v15.17L22 8v2.83L11 22Z"/>`),
		g.Group(children),
	)
}
package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardwareCircuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 0H2a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2ZM6 16a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1ZM22 6h-2.184a3 3 0 1 0 0 2H22v4h-4v2h4v2h-2v2h2v4h-8v-1.184a3 3 0 1 0-2 0V22H7v-4.184a3 3 0 1 0-2 0V22H2V2h4v6h2V2h2v10h2V2h10Zm-4 1a1 1 0 1 1-1-1a1.001 1.001 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}
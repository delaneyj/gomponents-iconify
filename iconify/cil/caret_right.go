package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M148.092 496h-36.45V16.333h34.547L416 256.286Zm-4.45-439.108v400.15l224.287-200.684Z"/>`),
		g.Group(children),
	)
}
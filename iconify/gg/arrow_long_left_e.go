package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongLeftE(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m5.263 7.757l-4.25 4.236l4.236 4.25l1.417-1.412l-1.815-1.82l12.117.04l-.008 2l6 .027l.026-6l-6-.027l-.009 2l-12.144-.04l1.842-1.837l-1.412-1.417Zm15.714 3.312l-.009 2l-2-.01l.01-2l2 .01Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
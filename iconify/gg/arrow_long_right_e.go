package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongRightE(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m22.986 11.993l-4.235 4.25l-1.417-1.412l1.815-1.82l-12.118.04l.01 2l-6 .027l-.028-6l6-.027l.01 2l12.144-.04l-1.842-1.837l1.412-1.417l4.25 4.236Zm-19.964-.924l.01 2l2-.01l-.01-2l-2 .01Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
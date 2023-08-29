package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackageUpgrade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.77 14.25L20 16.49V8h2V6H2v2h2v11a2 2 0 0 0 2 2h5ZM6 16h5v2H6Z"/><path fill="currentColor" d="m22.01 20.53l-2.83-2.83l-1.41-1.41l-4.25 4.25l1.41 1.41l2.84-2.83l2.83 2.83l1.41-1.42z"/>`),
		g.Group(children),
	)
}
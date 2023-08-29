package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeClimate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.495 8.11l-6.364 6.365l1.414 1.414l4.95-4.95l4.95 4.95l1.414-1.414L9.495 8.11Zm5.01 0l-1.973 1.974l1.418 1.41l.555-.555l4.95 4.95l1.414-1.414l-6.364-6.364Z"/>`),
		g.Group(children),
	)
}
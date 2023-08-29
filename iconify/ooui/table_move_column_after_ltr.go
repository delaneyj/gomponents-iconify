package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableMoveColumnAfterLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m16 10l-5-4v3H6v2h5v3z"/><path fill="currentColor" d="M0 2h20v16H0zm5 6v4h5v4h8V4h-8v4z"/>`),
		g.Group(children),
	)
}
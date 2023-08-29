package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableMoveColumnAfterRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m4 10l5-4v3h5v2H9v3z"/><path fill="currentColor" d="M0 2v16h20V2zm2 2h8v4h5v4h-5v4H2z"/>`),
		g.Group(children),
	)
}
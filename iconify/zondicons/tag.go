package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M0 10V2l2-2h8l10 10l-10 10L0 10zm4.5-4a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3z"/>`),
		g.Group(children),
	)
}
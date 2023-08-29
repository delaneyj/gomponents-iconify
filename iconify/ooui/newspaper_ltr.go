package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewspaperLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 2a2 2 0 0 0-2 2v12a1 1 0 0 1-1-1V5h-.5A1.5 1.5 0 0 0 0 6.5v10A1.5 1.5 0 0 0 1.5 18H18a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm1 2h11v4H6zm0 6h6v1H6zm0 2h6v1H6zm0 2h6v1H6zm7-4h4v5h-4z"/>`),
		g.Group(children),
	)
}
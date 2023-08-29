package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpecialPagesRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 0a2 2 0 0 1 2 2H6a2 2 0 0 0-2 2v12a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2z"/><path fill="currentColor" d="M7 20a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2H7zm6.5-3.5l.41-1.09L15 15l-1.09-.41l-.41-1.09l-.41 1.09L12 15l1.09.41l.41 1.09zm-2.982-.949l.952-2.561l2.53-.964l-2.53-.964l-.952-2.562l-.952 2.562l-2.53.964l2.53.964l.952 2.561zM14 10.5l.547-1.453L16 8.5l-1.453-.547L14 6.5l-.547 1.453L12 8.5l1.453.547L14 10.5z"/>`),
		g.Group(children),
	)
}
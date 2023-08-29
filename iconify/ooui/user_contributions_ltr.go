package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserContributionsLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<circle cx="15.5" cy="10.5" r="2.5" fill="currentColor"/><path fill="currentColor" d="M1 15h8v2H1Zm0-6h10v2H1Zm0-6h16v2H1Zm14.5 10.6c-3.3 0-4.5 1.6-4.5 2.7V18h9v-1.7c0-1-1.2-2.7-4.5-2.7z"/>`),
		g.Group(children),
	)
}
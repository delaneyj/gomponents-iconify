package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserContributionsRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<circle cx="4.5" cy="10.5" r="2.5" fill="currentColor"/><path fill="currentColor" d="M19 15h-8v2h8zm0-6H9v2h10Zm0-6H3v2h16ZM4.5 13.6c3.3 0 4.5 1.6 4.5 2.7V18H0v-1.7c0-1 1.2-2.7 4.5-2.7z"/>`),
		g.Group(children),
	)
}
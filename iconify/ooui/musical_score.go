package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicalScore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 15V5h8v10h2V2H6v13"/><circle cx="15" cy="15" r="3" fill="currentColor"/><circle cx="5" cy="15" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}
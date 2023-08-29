package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapTrail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m20 6l-1-1l-1.5 1.5L16 5l-1 1l1.5 1.5L15 9l1 1l1.5-1.5L19 10l1-1l-1.5-1.5z"/><circle cx="7.5" cy="14.5" r="3.5" fill="currentColor"/><circle cx="7" cy="3" r="2" fill="currentColor"/><circle cx="13" cy="7" r="1" fill="currentColor"/><circle cx="10" cy="6" r="1" fill="currentColor"/><circle cx="3" cy="3" r="1" fill="currentColor"/><circle cx="1" cy="6" r="1" fill="currentColor"/><circle cx="1" cy="9" r="1" fill="currentColor"/><circle cx="3" cy="12" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}
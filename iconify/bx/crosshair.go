package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crosshair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm1 17.931V17h-2v2.931A8.008 8.008 0 0 1 4.069 13H7v-2H4.069A8.008 8.008 0 0 1 11 4.069V7h2V4.069A8.007 8.007 0 0 1 19.931 11H17v2h2.931A8.008 8.008 0 0 1 13 19.931z"/>`),
		g.Group(children),
	)
}
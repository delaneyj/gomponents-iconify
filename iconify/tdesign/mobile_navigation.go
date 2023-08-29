package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileNavigation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 1h16v22H4V1Zm2 2v18h12V3H6Zm6 1.764l4.946 9.892L12 13.11l-4.946 1.546L12 4.764Zm-1.054 6.58l1.054-.33l1.054.33L12 9.236l-1.054 2.108ZM11 17h2.004v2.004H11V17Z"/>`),
		g.Group(children),
	)
}
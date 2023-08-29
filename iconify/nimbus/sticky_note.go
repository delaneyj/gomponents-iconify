package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.25 1.5H1.75A1.25 1.25 0 0 0 .5 2.75v10.5a1.25 1.25 0 0 0 1.25 1.25h8.69a1.24 1.24 0 0 0 .93-.41l3.81-4.23A1.24 1.24 0 0 0 15.5 9V2.75a1.25 1.25 0 0 0-1.25-1.25zM1.75 13.25V2.75h12.5v5h-3.81A1.25 1.25 0 0 0 9.19 9v4.23zm8.69 0V9h3.81z"/>`),
		g.Group(children),
	)
}
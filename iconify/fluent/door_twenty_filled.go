package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 4A1.5 1.5 0 0 1 6 2.5h8A1.5 1.5 0 0 1 15.5 4v12a1.5 1.5 0 0 1-1.5 1.5H6A1.5 1.5 0 0 1 4.5 16V4ZM7 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}
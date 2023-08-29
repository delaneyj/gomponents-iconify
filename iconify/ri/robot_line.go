package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 4.055A9 9 0 0 1 21 13v9H3v-9a9 9 0 0 1 8-8.945V1h2v3.055ZM19 20v-7a7 7 0 1 0-14 0v7h14Zm-7-2a5 5 0 1 1 0-10a5 5 0 0 1 0 10Zm0-2a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-2a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}
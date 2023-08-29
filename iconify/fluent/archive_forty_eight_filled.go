package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.25 8.75a2.5 2.5 0 0 1 2.5-2.5h30.5a2.5 2.5 0 0 1 2.5 2.5v4a2.5 2.5 0 0 1-2.5 2.5H8.75a2.5 2.5 0 0 1-2.5-2.5v-4Zm2 9h31.5V34.5a7.25 7.25 0 0 1-7.25 7.25h-17a7.25 7.25 0 0 1-7.25-7.25V17.75Zm12.25 4.5a1.25 1.25 0 1 0 0 2.5h7a1.25 1.25 0 1 0 0-2.5h-7Z"/>`),
		g.Group(children),
	)
}
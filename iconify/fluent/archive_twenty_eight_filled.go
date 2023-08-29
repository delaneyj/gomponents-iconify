package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 4.5A1.5 1.5 0 0 1 4.5 3h19A1.5 1.5 0 0 1 25 4.5v2A1.5 1.5 0 0 1 23.5 8h-19A1.5 1.5 0 0 1 3 6.5v-2Zm1 5h20v10.75A4.75 4.75 0 0 1 19.25 25H8.75A4.75 4.75 0 0 1 4 20.25V9.5Zm7.75 3a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5h-4.5Z"/>`),
		g.Group(children),
	)
}
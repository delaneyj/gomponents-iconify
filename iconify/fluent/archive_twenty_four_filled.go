package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 4.5A1.5 1.5 0 0 1 4.5 3h15A1.5 1.5 0 0 1 21 4.5V6a1.5 1.5 0 0 1-1.5 1.5h-15A1.5 1.5 0 0 1 3 6V4.5ZM4 9h16v8.25A3.75 3.75 0 0 1 16.25 21h-8.5A3.75 3.75 0 0 1 4 17.25V9Zm6.25 2a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-3.5Z"/>`),
		g.Group(children),
	)
}
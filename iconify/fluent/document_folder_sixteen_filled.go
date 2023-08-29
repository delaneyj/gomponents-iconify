package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentFolderSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.5 9H12V4.085A1.5 1.5 0 0 1 13 5.5v3.55a2.512 2.512 0 0 0-.5-.05ZM6.854 4.732L11 8.88V3.5A1.5 1.5 0 0 0 9.5 2h-5A1.5 1.5 0 0 0 3 3.5v.55c.162-.033.329-.05.5-.05h1.586a2.5 2.5 0 0 1 1.768.732ZM3.5 5A1.5 1.5 0 0 0 2 6.5V12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-.5a1.5 1.5 0 0 0-1.5-1.5h-1.586a.5.5 0 0 1-.353-.146L6.146 5.439A1.5 1.5 0 0 0 5.086 5H3.5Z"/>`),
		g.Group(children),
	)
}
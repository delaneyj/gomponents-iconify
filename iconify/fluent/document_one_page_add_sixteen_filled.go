package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentOnePageAddSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 1a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h2.257A5.5 5.5 0 0 1 13 6.207V3a2 2 0 0 0-2-2H5Zm4.59 3.992L9.5 5h-3a.5.5 0 0 1-.09-.992L6.5 4h3a.5.5 0 0 1 .09.992ZM16 11.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-4-2a.5.5 0 0 0-1 0V11H9.5a.5.5 0 0 0 0 1H11v1.5a.5.5 0 0 0 1 0V12h1.5a.5.5 0 0 0 0-1H12V9.5Z"/>`),
		g.Group(children),
	)
}
package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 1h16v10h-2V3H6v18h6.5v2H4V1Zm13.25 13.5a2.75 2.75 0 1 1 0 5.5a2.75 2.75 0 0 1 0-5.5Zm3.992 5.325a4.75 4.75 0 1 0-1.414 1.415l1.67 1.674l1.416-1.412l-1.672-1.677Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeTextTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.5 5a.5.5 0 0 0 0 1H8a.5.5 0 0 0 0-1H2.5Zm3 3a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5ZM4 11.5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1h-10a.5.5 0 0 1-.5-.5ZM2.5 14a.5.5 0 0 0 0 1H11a.5.5 0 0 0 0-1H2.5Zm10-5.5A.5.5 0 0 1 13 8h3.5a.5.5 0 0 1 0 1H13a.5.5 0 0 1-.5-.5Zm-2-3.5a.5.5 0 0 0 0 1h7a.5.5 0 0 0 0-1h-7Z"/>`),
		g.Group(children),
	)
}
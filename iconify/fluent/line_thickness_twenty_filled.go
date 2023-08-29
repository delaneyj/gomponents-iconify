package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineThicknessTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 4a.75.75 0 0 0 0 1.5h14.5a.75.75 0 0 0 0-1.5H2.75ZM2 9.25C2 8.56 2.56 8 3.25 8h13.5a1.25 1.25 0 1 1 0 2.5H3.25C2.56 10.5 2 9.94 2 9.25Zm0 5.5c0-.966.784-1.75 1.75-1.75h12.5a1.75 1.75 0 1 1 0 3.5H3.75A1.75 1.75 0 0 1 2 14.75Z"/>`),
		g.Group(children),
	)
}
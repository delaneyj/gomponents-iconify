package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineThicknessTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 3.5a.75.75 0 0 0 0 1.5h18.5a.75.75 0 0 0 0-1.5H2.75ZM2 10.75c0-.69.56-1.25 1.25-1.25h17.5a1.25 1.25 0 1 1 0 2.5H3.25C2.56 12 2 11.44 2 10.75Zm0 7.5c0-.966.784-1.75 1.75-1.75h16.5a1.75 1.75 0 1 1 0 3.5H3.75A1.75 1.75 0 0 1 2 18.25Z"/>`),
		g.Group(children),
	)
}
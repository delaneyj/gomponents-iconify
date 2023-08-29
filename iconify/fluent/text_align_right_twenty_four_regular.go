package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignRightTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 5.75A.75.75 0 0 1 5.75 5h15.5a.75.75 0 0 1 0 1.5H5.75A.75.75 0 0 1 5 5.75Zm5 13a.75.75 0 0 1 .75-.75h10.5a.75.75 0 0 1 0 1.5h-10.5a.75.75 0 0 1-.75-.75ZM2.75 11.5a.75.75 0 0 0 0 1.5h18.5a.75.75 0 0 0 0-1.5H2.75Z"/>`),
		g.Group(children),
	)
}
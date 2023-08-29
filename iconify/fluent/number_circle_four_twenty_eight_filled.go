package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFourTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 14C2 7.373 7.373 2 14 2s12 5.373 12 12s-5.373 12-12 12S2 20.627 2 14Zm14.5-5.581c0-1.092-1.42-1.517-2.02-.605l-5.403 8.214a.95.95 0 0 0 .794 1.472h5.128v1.75a.75.75 0 0 0 1.5 0V17.5H17.5a.75.75 0 0 0 0-1.5h-1V8.419ZM15 9.755V16h-4.11L15 9.755Z"/>`),
		g.Group(children),
	)
}
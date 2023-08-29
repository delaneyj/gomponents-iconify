package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentPageBreakTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.5 12a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5A.75.75 0 0 1 2.5 12Zm4 0a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5A.75.75 0 0 1 6.5 12Zm4 0a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5a.75.75 0 0 1-.75-.75Zm4 0a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5a.75.75 0 0 1-.75-.75Zm4 0a.75.75 0 0 1 .75-.75h1.5a.75.75 0 0 1 0 1.5h-1.5a.75.75 0 0 1-.75-.75ZM4.75 2a.75.75 0 0 0-.75.75V7a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V2.75a.75.75 0 0 0-1.5 0V7a.5.5 0 0 1-.5.5H6a.5.5 0 0 1-.5-.5V2.75A.75.75 0 0 0 4.75 2Zm14.5 20a.75.75 0 0 0 .75-.75V17a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v4.25a.75.75 0 0 0 1.5 0V17a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 .5.5v4.25c0 .414.336.75.75.75Z"/>`),
		g.Group(children),
	)
}
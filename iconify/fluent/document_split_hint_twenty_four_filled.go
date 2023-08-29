package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentSplitHintTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 8V2H6a2 2 0 0 0-2 2v7.5a.5.5 0 0 0 .5.5h15a.5.5 0 0 0 .5-.5V10h-6a2 2 0 0 1-2-2Zm1.5 0V2.5l6 6H14a.5.5 0 0 1-.5-.5Zm-8 5.75a.75.75 0 0 0-1.5 0v2.495a.75.75 0 0 0 1.5 0V13.75ZM19.25 13a.75.75 0 0 1 .75.75v2.495a.75.75 0 0 1-1.5 0V13.75a.75.75 0 0 1 .75-.75Zm-9 7.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 0-1.5h-3.5Zm-5.5-2a.75.75 0 0 1 .75.75V20a.5.5 0 0 0 .5.5h1.25a.75.75 0 0 1 0 1.5H6a2 2 0 0 1-2-2v-.75a.75.75 0 0 1 .75-.75Zm13.75.75a.75.75 0 0 1 1.5 0V20a2 2 0 0 1-2 2h-1.25a.75.75 0 0 1 0-1.5H18a.5.5 0 0 0 .5-.5v-.75Z"/>`),
		g.Group(children),
	)
}
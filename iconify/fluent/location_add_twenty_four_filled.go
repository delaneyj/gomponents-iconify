package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationAddTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m16.97 18.056l1.187-1.174a8.707 8.707 0 1 0-12.314 0c.296.296.663.659 1.102 1.09l3.491 3.396a2.25 2.25 0 0 0 3.128 0a657.733 657.733 0 0 0 3.406-3.312ZM12 7a.75.75 0 0 1 .75.75V10h2.5a.75.75 0 0 1 0 1.5h-2.5v2.25a.75.75 0 0 1-1.5 0V11.5h-2.5a.75.75 0 0 1 0-1.5h2.5V7.75A.75.75 0 0 1 12 7Z"/>`),
		g.Group(children),
	)
}
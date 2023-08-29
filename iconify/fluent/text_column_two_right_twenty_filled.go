package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColumnTwoRightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 5.75a.75.75 0 0 0-.75-.75h-2.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 .75-.75Zm-6 0a.75.75 0 0 0-.75-.75h-6.5a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 .75-.75Zm6 3a.75.75 0 0 0-.75-.75h-2.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 .75-.75Zm-6 0a.75.75 0 0 0-.75-.75h-6.5a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 .75-.75Zm6 3a.75.75 0 0 0-.75-.75h-2.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 .75-.75Zm-6 0a.75.75 0 0 0-.75-.75h-6.5a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 .75-.75Zm6 3a.75.75 0 0 0-.75-.75h-2.5a.75.75 0 0 0 0 1.5h2.5a.75.75 0 0 0 .75-.75Zm-6 0a.75.75 0 0 0-.75-.75h-6.5a.75.75 0 0 0 0 1.5h6.5a.75.75 0 0 0 .75-.75Z"/>`),
		g.Group(children),
	)
}
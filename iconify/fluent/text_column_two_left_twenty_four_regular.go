package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColumnTwoLeftTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 5.75a.75.75 0 0 0-.75-.75h-9.5a.75.75 0 0 0 0 1.5h9.5a.75.75 0 0 0 .75-.75Zm-13 0A.75.75 0 0 0 7.25 5h-3.5a.75.75 0 0 0 0 1.5h3.5A.75.75 0 0 0 8 5.75Zm13 4a.75.75 0 0 0-.75-.75h-9.5a.75.75 0 0 0 0 1.5h9.5a.75.75 0 0 0 .75-.75Zm-13 0A.75.75 0 0 0 7.25 9h-3.5a.75.75 0 0 0 0 1.5h3.5A.75.75 0 0 0 8 9.75Zm13 4a.75.75 0 0 0-.75-.75h-9.5a.75.75 0 0 0 0 1.5h9.5a.75.75 0 0 0 .75-.75Zm-13 0a.75.75 0 0 0-.75-.75h-3.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 .75-.75Zm13 4a.75.75 0 0 0-.75-.75h-9.5a.75.75 0 0 0 0 1.5h9.5a.75.75 0 0 0 .75-.75Zm-13 0a.75.75 0 0 0-.75-.75h-3.5a.75.75 0 0 0 0 1.5h3.5a.75.75 0 0 0 .75-.75Z"/>`),
		g.Group(children),
	)
}
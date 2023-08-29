package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardLayoutSplitTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.745 5a2.25 2.25 0 0 1 2.25 2.25v9.505a2.25 2.25 0 0 1-2.25 2.25H4.25A2.25 2.25 0 0 1 2 16.755V7.25A2.25 2.25 0 0 1 4.25 5h15.495ZM6.5 16H10a.75.75 0 0 0 0-1.5H6.5a.75.75 0 0 0 0 1.5Zm7.5 0h3.5a.75.75 0 0 0 0-1.5H14a.75.75 0 0 0 0 1.5Zm3.5-5a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-2.995 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-5 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM6 8a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm2.995 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm6 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm3 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}
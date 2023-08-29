package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 2.75A.75.75 0 0 0 8.25 2h-5.5a.75.75 0 0 0-.75.75v5.5a.75.75 0 0 0 1.5 0V4.56l9.22 9.22a.75.75 0 0 0 1.06-1.06L4.56 3.5h3.69A.75.75 0 0 0 9 2.75Z"/>`),
		g.Group(children),
	)
}
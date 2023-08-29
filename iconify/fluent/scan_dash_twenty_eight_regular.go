package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanDashTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.75 3A3.75 3.75 0 0 0 3 6.75v3.5a.75.75 0 0 0 1.5 0v-3.5A2.25 2.25 0 0 1 6.75 4.5h3.501a.75.75 0 0 0 0-1.5H6.75Zm11 0a.75.75 0 0 0 0 1.5h3.5a2.25 2.25 0 0 1 2.25 2.25v3.5a.75.75 0 0 0 1.5 0v-3.5A3.75 3.75 0 0 0 21.25 3h-3.5ZM4.5 17.75a.75.75 0 0 0-1.5 0v3.5A3.75 3.75 0 0 0 6.75 25h3.5a.75.75 0 0 0 0-1.5h-3.5a2.25 2.25 0 0 1-2.25-2.25v-3.5Zm20.5 0a.75.75 0 0 0-1.5 0v3.5a2.25 2.25 0 0 1-2.25 2.25h-3.5a.75.75 0 0 0 0 1.5h3.5A3.75 3.75 0 0 0 25 21.25v-3.5Zm-16.75-4.5a.75.75 0 0 0 0 1.5h11.5a.75.75 0 0 0 0-1.5H8.25Z"/>`),
		g.Group(children),
	)
}
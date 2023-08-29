package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.37 1.222a1 1 0 0 1 1.26 0l3.814 3.09A1.5 1.5 0 0 1 11 5.476V10a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V7H5v3a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V5.477a1.5 1.5 0 0 1 .556-1.166l3.815-3.089Zm4.445 3.866L6 2L2.185 5.088A.5.5 0 0 0 2 5.477V10h2V7a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v3h2V5.477a.5.5 0 0 0-.185-.389Z"/>`),
		g.Group(children),
	)
}
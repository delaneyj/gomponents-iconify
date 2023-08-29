package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanDashThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.75 3A4.75 4.75 0 0 0 3 7.75v3.354a1.25 1.25 0 0 0 2.5 0V7.75A2.25 2.25 0 0 1 7.75 5.5h3.354a1.25 1.25 0 0 0 0-2.5H7.75Zm13.146 0a1.25 1.25 0 0 0 0 2.5h3.354a2.25 2.25 0 0 1 2.25 2.25v3.354a1.25 1.25 0 0 0 2.5 0V7.75A4.75 4.75 0 0 0 24.25 3h-3.354ZM5.5 20.896a1.25 1.25 0 0 0-2.5 0v3.354A4.75 4.75 0 0 0 7.75 29h3.354a1.25 1.25 0 0 0 0-2.5H7.75a2.25 2.25 0 0 1-2.25-2.25v-3.354Zm23.5 0a1.25 1.25 0 0 0-2.5 0v3.354a2.25 2.25 0 0 1-2.25 2.25h-3.354a1.25 1.25 0 0 0 0 2.5h3.354A4.75 4.75 0 0 0 29 24.25v-3.354ZM9.5 14.75a1.25 1.25 0 1 0 0 2.5h13a1.25 1.25 0 1 0 0-2.5h-13Z"/>`),
		g.Group(children),
	)
}
package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M13 16.5a11.003 11.003 0 0 0 4.041 3.93l1.459.82v.25h-13v-.25l1.459-.82A11.003 11.003 0 0 0 11 16.5m-9.5-3h21m0-11h-21v14h21v-14Z"/>`),
		g.Group(children),
	)
}
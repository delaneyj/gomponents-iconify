package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EverydayTasks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.272 18.654A20 20 0 0 1 44 24h0a20 20 0 1 1-8.678-16.487"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.873 24.204l7.295 7.376l22.34-23.092"/>`),
		g.Group(children),
	)
}
package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopwatchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.5 5v3.5H10m-4-8h3m-1.5 2a6 6 0 1 0 0 12a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}
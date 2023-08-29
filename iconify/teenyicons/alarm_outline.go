package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5 3.5l3-3m8 0l3 3M7.5 5v3.5H10m-2.5-6a6 6 0 1 0 0 12a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}
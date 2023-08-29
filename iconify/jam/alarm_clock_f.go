package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmClockF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.962 2.05a5 5 0 1 1 7.668 6.387a9.008 9.008 0 0 0-7.668-6.386zM1.37 8.439a5 5 0 1 1 7.668-6.387A9.008 9.008 0 0 0 1.37 8.438zM10 18a7 7 0 1 1 0-14a7 7 0 0 1 0 14zm1-8V6a1 1 0 0 0-2 0v5a1 1 0 0 0 1 1h3a1 1 0 0 0 0-2h-2z"/>`),
		g.Group(children),
	)
}
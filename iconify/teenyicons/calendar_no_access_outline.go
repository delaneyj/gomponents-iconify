package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarNoAccessOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 0v5m8-5v5m-2 1.5l-4 4m-4-8h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1Zm6 9a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		g.Group(children),
	)
}
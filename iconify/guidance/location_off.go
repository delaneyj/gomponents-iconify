package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m16.19 16.19l1.337-1.215A8.365 8.365 0 0 0 20.25 8.8c0-4.558-3.694-8.3-8.25-8.3c-3.104 0-5.807 1.737-7.216 4.284M12 23.92a9.041 9.041 0 0 1 2.383-6.037M12 23.92a9.04 9.04 0 0 0-2.96-6.61l-2.567-2.334a8.365 8.365 0 0 1-2.6-7.603M12 23.92V24v-.057M.5.5l23 23M9.564 9.564a2.5 2.5 0 1 1 1.872 1.872"/>`),
		g.Group(children),
	)
}
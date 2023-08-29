package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Block(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zM4 12a8 8 0 0 1 8-8c1.848 0 3.545.633 4.9 1.686L5.686 16.9A7.957 7.957 0 0 1 4 12zm8 8a7.955 7.955 0 0 1-4.9-1.686L18.314 7.1A7.957 7.957 0 0 1 20 12a8 8 0 0 1-8 8z"/>`),
		g.Group(children),
	)
}
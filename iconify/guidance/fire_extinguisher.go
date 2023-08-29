package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireExtinguisher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M15.5 3A1.5 1.5 0 0 1 17 1.5h3.5v3H17A1.5 1.5 0 0 1 15.5 3Zm0 0h-3a8 8 0 0 0-8 8v4A2.5 2.5 0 0 1 2 17.5M12 3V0m4.5 12.5H12v5h4.5m0-7v13h-9v-13a4.5 4.5 0 0 1 9 0Z"/>`),
		g.Group(children),
	)
}
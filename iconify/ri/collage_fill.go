package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollageFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.19 13.264l1.382 7.841l-8.571.001a1 1 0 0 1-1-1L3 14.708l8.19-1.444ZM20 3.106a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-5.397l-3.174-18h8.572Zm-10.602 0l1.444 8.188L3 12.677v-8.57a1 1 0 0 1 1-1l5.398-.002Z"/>`),
		g.Group(children),
	)
}
package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdminPost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10.44 3.02l1.82-1.82l6.36 6.35l-1.83 1.82a2.731 2.731 0 0 0-3.41.36l-.75.75c-.92.93-1.04 2.35-.35 3.41l-1.83 1.82l-2.41-2.41l-2.8 2.79c-.42.42-3.38 2.71-3.8 2.29s1.86-3.39 2.28-3.81l2.79-2.79L4.1 9.36l1.83-1.82c1.05.69 2.48.57 3.4-.36l.75-.75c.93-.92 1.05-2.35.36-3.41z"/>`),
		g.Group(children),
	)
}
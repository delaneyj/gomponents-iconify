package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Graph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 648h717v-56H56V46H0v602zm98-110h616V96L584 215l-121-44l-128 146l-131-47L98 380v158z"/>`),
		g.Group(children),
	)
}
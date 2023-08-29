package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Target(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0c6.627 0 12 5.373 12 12s-5.373 12-12 12S0 18.627 0 12S5.373 0 12 0zm0 19.826a7.827 7.827 0 1 0 0-15.652a7.827 7.827 0 0 0 0 15.652zm0-3.985a3.84 3.84 0 1 1 0-7.68a3.84 3.84 0 0 1 0 7.68z"/>`),
		g.Group(children),
	)
}
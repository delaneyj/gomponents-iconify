package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Album(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M0 0h20v20H0V0zm10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16zm0-5a3 3 0 1 1 0-6a3 3 0 0 1 0 6z"/>`),
		g.Group(children),
	)
}
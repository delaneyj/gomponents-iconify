package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineKitchen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 9V4c0-1.1-.9-2-2-2H6c-1.1 0-2 .9-2 2v5h16zM8 5h2v3H8V5zm-4 6v9c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2v-9H4zm6 6H8v-5h2v5z"/>`),
		g.Group(children),
	)
}
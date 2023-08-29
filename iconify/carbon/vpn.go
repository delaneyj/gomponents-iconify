package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 23h-2V9h6a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-4zm0-7h4v-5h-4zm14 3L24.32 9H22v14h2V13l3.68 10H30V9h-2v10zM8 9L6 22L4 9H2l2.52 14h2.96L10 9H8z"/>`),
		g.Group(children),
	)
}
package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneViewComfy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 18h10v-5H10v5zM4 6v5h16V6H4zm0 12h4v-5H4v5z" opacity=".3"/><path fill="currentColor" d="M2 4v16h20V4H2zm6 14H4v-5h4v5zm12 0H10v-5h10v5zm0-7H4V6h16v5z"/>`),
		g.Group(children),
	)
}
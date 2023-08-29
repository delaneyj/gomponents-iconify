package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0h27.65v5.219H0zm0 9.39h27.65v5.219H0zm0 9.391h27.65V24H0z"/>`),
		g.Group(children),
	)
}
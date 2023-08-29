package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 6v2h4V6zm12 0v2h4V6zM4 9v18h24V9zm2 2h20v14H6zm14 4v2h-2v2h2v2h2v-2h2v-2h-2v-2zM8 17v2h6v-2z"/>`),
		g.Group(children),
	)
}
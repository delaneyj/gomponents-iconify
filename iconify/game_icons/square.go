package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Square(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M23.05 23.05V488.9H488.9V23.05zm17.9 17.9H471.1V471.1H40.95z"/>`),
		g.Group(children),
	)
}
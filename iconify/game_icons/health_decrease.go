package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthDecrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M196 16c-16.62 0-30 13.38-30 30v120H46c-16.62 0-30 13.38-30 30v120c0 16.62 13.38 30 30 30h120v120c0 16.62 13.38 30 30 30h120c16.62 0 30-13.38 30-30V346h120c16.62 0 30-13.38 30-30V196c0-16.62-13.38-30-30-30H346V46c0-16.62-13.38-30-30-30zm0 135h120l-60 60zm0 75h120l-60 60zm0 75h120l-60 60z"/>`),
		g.Group(children),
	)
}
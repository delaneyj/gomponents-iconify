package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionSigns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m277.3 35.11l-32.1 2.12l-1 32.93l33.6.05l-.5-35.1zM151.7 88.04L67.96 118.9L152 169.2l188-6.2l2.4-74.72l-190.7-.24zM279.2 183l-38.3 1.3l-.8 27.8l39.5 1.2l-.4-30.3zm-93.1 45.6l.8 64.2l200.4.9l51.7-29.2l-55.8-30.2l-197.1-5.7zm51.1 82.5l-5 175.3l50.9.6l-2.3-175.7l-43.6-.2z"/>`),
		g.Group(children),
	)
}
package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libya(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m477.6 99.68l8.5 363.72L454 465L211.7 351.9l-48.1 28.2l-137.68-94l14.52-133L98.59 46.98L199.4 77.54l13.1 25.96l97 39.7l34.4-13.8l-19.1-45.75l82.5-22.92z"/>`),
		g.Group(children),
	)
}
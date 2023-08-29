package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m290.352 13.213l-11.475 218.984l204.68-78.584a219.242 219.284 0 0 0-193.205-140.4zm-51.39 47.566A219.242 219.284 0 0 0 38.59 206.24a219.242 219.284 0 0 0 77.3 250.918a219.242 219.284 0 0 0 262.49-3.092a219.242 219.284 0 0 0 71.366-252.67l-204.682 78.583l12.24-218.943a219.242 219.284 0 0 0-18.34-.258z"/>`),
		g.Group(children),
	)
}
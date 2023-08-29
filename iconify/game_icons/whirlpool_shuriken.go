package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhirlpoolShuriken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m286.077 121.043l45.124-103.484l-93.427 69.007l-.392-.282l-.062.616l-105.5 77.925l-83.85-76.773l35.532 112.459l-.114.082l.162.071l39.185 124.029l-98.881 60.209l249.358-5.869L298.8 494.441l75.664-240.184l113.682 10.997zm-46.453 177.914c-32.097 0-58.115-26.019-58.115-58.114c0-32.098 26.018-58.115 58.115-58.115s58.114 26.017 58.114 58.115c0 32.096-26.017 58.114-58.114 58.114z"/>`),
		g.Group(children),
	)
}
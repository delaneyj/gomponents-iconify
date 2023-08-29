package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnticlockwiseRotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M248.91 50a205.9 205.9 0 0 1 35.857 3.13c85.207 15.025 152.077 81.895 167.102 167.102c15.023 85.208-24.944 170.917-99.874 214.178c-32.782 18.927-69.254 27.996-105.463 27.553c-46.555-.57-92.675-16.865-129.957-48.15l30.855-36.768a157.846 157.846 0 0 0 180.566 15.797a157.846 157.846 0 0 0 76.603-164.274A157.848 157.848 0 0 0 276.429 100.4a157.84 157.84 0 0 0-139.17 43.862L185 192H57V64l46.34 46.342C141.758 71.962 194.17 50.03 248.91 50z"/>`),
		g.Group(children),
	)
}
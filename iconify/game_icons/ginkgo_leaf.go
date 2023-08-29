package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GinkgoLeaf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m118.9 20.13l-16.4 7.56c21.9 47.21 55.3 102.71 90.7 152.61C127.8 322 18.79 349.2 29.06 377.1C80.7 517.8 239.3 503.5 314.3 467.7c27.5-13.1 5.1-61.3-6-94c20.6 25.8 61.5 67.9 77.1 58.6c95.9-57.3 119-164 80.5-267.8c-20.2-54.4-116.3 31.7-257.6 5.8c-35-49.2-68.2-104.28-89.4-150.17z"/>`),
		g.Group(children),
	)
}
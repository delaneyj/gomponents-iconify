package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uruguay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m86.627 56.843l86.84-33.014c98.7 82.092 216.765 136.019 295.685 246.884c-65.653 71.335-17.208 71.745-20.095 104.064C426.444 500.632 315.248 482.712 225.14 488.17c-47.279-23.845-88.53-54.013-128.466-46.65L42.848 392z"/>`),
		g.Group(children),
	)
}
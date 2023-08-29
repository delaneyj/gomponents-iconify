package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineNavigation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 7.27l4.28 10.43l-3.47-1.53l-.81-.36l-.81.36l-3.47 1.53L12 7.27M12 2L4.5 20.29l.71.71L12 18l6.79 3l.71-.71L12 2z"/>`),
		g.Group(children),
	)
}
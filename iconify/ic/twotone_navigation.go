package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneNavigation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.72 17.7l3.47-1.53l.81-.36l.81.36l3.47 1.53L12 7.27z" opacity=".3"/><path fill="currentColor" d="m4.5 20.29l.71.71L12 18l6.79 3l.71-.71L12 2L4.5 20.29zm8.31-4.12l-.81-.36l-.81.36l-3.47 1.53L12 7.27l4.28 10.43l-3.47-1.53z"/>`),
		g.Group(children),
	)
}
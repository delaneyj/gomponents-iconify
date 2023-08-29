package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineFlashOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.27 3L2 4.27l5 5V13h3v9l3.58-6.14L17.73 20L19 18.73L3.27 3zM17 10h-4l4-8H7v2.18l8.46 8.46L17 10z"/>`),
		g.Group(children),
	)
}
package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkMinimalisticTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m14.162 18.488l-.72.72a6.117 6.117 0 0 1-8.65-8.65l.72-.72m4.325 4.324l4.325-4.325M9.837 5.512l.721-.72a6.117 6.117 0 0 1 8.65 0m-.72 9.37l.72-.72A6.099 6.099 0 0 0 20.998 9"/>`),
		g.Group(children),
	)
}
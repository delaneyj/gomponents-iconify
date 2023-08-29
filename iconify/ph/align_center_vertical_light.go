package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterVerticalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 122h-18V72a14 14 0 0 0-14-14h-40a14 14 0 0 0-14 14v50h-20V48a14 14 0 0 0-14-14H64a14 14 0 0 0-14 14v74H32a6 6 0 0 0 0 12h18v74a14 14 0 0 0 14 14h40a14 14 0 0 0 14-14v-74h20v50a14 14 0 0 0 14 14h40a14 14 0 0 0 14-14v-50h18a6 6 0 0 0 0-12Zm-118 86a2 2 0 0 1-2 2H64a2 2 0 0 1-2-2V48a2 2 0 0 1 2-2h40a2 2 0 0 1 2 2Zm88-24a2 2 0 0 1-2 2h-40a2 2 0 0 1-2-2V72a2 2 0 0 1 2-2h40a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterVerticalSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 122h-34V48a14 14 0 0 0-14-14H96a14 14 0 0 0-14 14v74H48a6 6 0 0 0 0 12h34v74a14 14 0 0 0 14 14h64a14 14 0 0 0 14-14v-74h34a6 6 0 0 0 0-12Zm-46 86a2 2 0 0 1-2 2H96a2 2 0 0 1-2-2V48a2 2 0 0 1 2-2h64a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}
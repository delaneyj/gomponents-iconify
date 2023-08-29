package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VipCrownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.005 19h20v2h-20v-2Zm0-14l5 3l5-6l5 6l5-3v12h-20V5Z"/>`),
		g.Group(children),
	)
}
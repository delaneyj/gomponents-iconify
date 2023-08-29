package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitVerticalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 152a12 12 0 0 1-12 12h-68v39l11.51-11.52a12 12 0 0 1 17 17l-32 32a12 12 0 0 1-17 0l-32-32a12 12 0 0 1 17-17L116 203v-39H48a12 12 0 0 1 0-24h160a12 12 0 0 1 12 12ZM48 116h160a12 12 0 0 0 0-24h-68V53l11.51 11.52a12 12 0 1 0 17-17l-32-32a12 12 0 0 0-17 0l-32 32a12 12 0 1 0 17 17L116 53v39H48a12 12 0 0 0 0 24Z"/>`),
		g.Group(children),
	)
}
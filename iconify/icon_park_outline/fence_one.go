package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FenceOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 4v40m36-32H4m36 32L8 12m0 32l32-32M27 44L8 25m32 6L21 12M8 31l18-19m-5 32l19-19m4 19H4M40 4v40"/>`),
		g.Group(children),
	)
}
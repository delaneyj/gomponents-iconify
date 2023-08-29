package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VpbankNeo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 16.417V40.25m0-23.833l8.667-8.667M24 16.417L15.333 7.75M4.5 19.667h10.833L24 28.333l8.667-8.666H43.5S27.25 40.25 25.083 40.25h-2.166C20.75 40.25 4.5 19.667 4.5 19.667"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.333 19.667V7.75L24 16.417l8.667-8.667v11.917M24 16.417V40.25"/>`),
		g.Group(children),
	)
}
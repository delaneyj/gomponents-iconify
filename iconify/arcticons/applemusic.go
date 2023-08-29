package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Applemusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.08 18.25v-7.83a2.16 2.16 0 0 1 1.75-2.13L38 4.54a2.16 2.16 0 0 1 2.58 2.13v7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.08 19.57a2.16 2.16 0 0 1 1.75-2.13L38 13.69a2.16 2.16 0 0 1 2.58 2.12m-.03-2.16v18.56M17.08 18.25v20.43"/><circle cx="12.27" cy="38.68" r="4.82" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="35.73" cy="32.21" r="4.82" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}
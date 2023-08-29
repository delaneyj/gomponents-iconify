package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5v16M19 5v9M7.641 3.645A5 5 0 0 1 12 5a5 5 0 0 0 7 0M5 14a5 5 0 0 1 7 0a4.984 4.984 0 0 0 3.437 1.429m3.019-.966c.19-.14.371-.294.544-.463M3 3l18 18"/>`),
		g.Group(children),
	)
}
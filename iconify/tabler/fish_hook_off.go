package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FishHookOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 9v3m-.085 3.924A5 5 0 0 1 6 15v-4l3 3m5-7a2 2 0 1 0 4 0a2 2 0 1 0-4 0m2-2V3M3 3l18 18"/>`),
		g.Group(children),
	)
}
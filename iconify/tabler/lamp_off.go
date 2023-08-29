package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LampOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20h6m-3 0v-8M7.325 7.35L5 12h7m4 0h3l-4-8H9l-.338.676M3 3l18 18"/>`),
		g.Group(children),
	)
}
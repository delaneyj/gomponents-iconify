package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VersionsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.184 6.162A2 2 0 0 1 12 5h6a2 2 0 0 1 2 2v9m-1.185 2.827A1.993 1.993 0 0 1 18 19h-6a2 2 0 0 1-2-2v-7M7 7v10M4 8v8M3 3l18 18"/>`),
		g.Group(children),
	)
}
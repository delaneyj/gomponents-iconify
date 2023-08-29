package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntennaOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M12 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="m7 23l1.111-4M17 23l-1.111-4M9.5 14l-1.389 5M9.5 14h4m-4 0l.8-2.88M8.11 19h7.778m0 0l-1.184-4.264M11.444 7L12 5l1.047 3.768M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
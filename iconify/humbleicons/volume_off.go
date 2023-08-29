package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M7 9H4v6h3l5 4V5L7 9z"/><path stroke-linecap="round" d="m16 9.5l5 5m0-5l-5 5"/></g>`),
		g.Group(children),
	)
}
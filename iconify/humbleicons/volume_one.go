package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M8 9H5v6h3l5 4V5L8 9z"/><path stroke-linecap="round" d="M17 8a5.657 5.657 0 0 1 0 8"/></g>`),
		g.Group(children),
	)
}
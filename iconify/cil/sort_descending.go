package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortDescending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M184 288h120v32H184zm0-72h184v32H184zm0-72h248v32H184zm0-72h312v32H184z"/><path fill="currentColor" d="M95.196 16v417.568l-51.883-51.882l-22.626 22.627l90.509 90.51l90.509-90.51l-22.627-22.626l-51.882 51.881V16h-32z"/>`),
		g.Group(children),
	)
}
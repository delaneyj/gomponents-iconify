package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Package(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 8l8 4M4 8v8l8 4M4 8l4-2m4 6l4-2m-4 2v8m8-12l-8-4l-4 2m12 2v8l-8 4m8-12l-4 2m0 0L8 6"/>`),
		g.Group(children),
	)
}
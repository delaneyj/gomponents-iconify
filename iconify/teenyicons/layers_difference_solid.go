package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersDifferenceSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4V1.5A1.5 1.5 0 0 1 5.5 0h8A1.5 1.5 0 0 1 15 1.5v8a1.5 1.5 0 0 1-1.5 1.5H11v2.5A1.5 1.5 0 0 1 9.5 15h-8A1.5 1.5 0 0 1 0 13.5v-8A1.5 1.5 0 0 1 1.5 4H4Zm6 1v5H5V5h5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
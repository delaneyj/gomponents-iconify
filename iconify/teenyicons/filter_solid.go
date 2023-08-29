package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 3H0V2h15v1Zm-3 5H3V7h9v1Zm-2 5H5v-1h5v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
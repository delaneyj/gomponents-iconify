package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuAltThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20 5V2H0v3h20zm0 6V8H0v3h20zm0 6v-3H0v3h20z"/>`),
		g.Group(children),
	)
}
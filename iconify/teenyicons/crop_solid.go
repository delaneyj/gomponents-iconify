package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 3V0h1v3h8v8h3v1h-3v3h-1v-3H3V6h1v5h7V4H0V3h3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
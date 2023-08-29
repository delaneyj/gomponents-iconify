package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PetsAllowed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M16 22V7.5h1A2.5 2.5 0 0 0 19.5 5V3.5H14A2.5 2.5 0 0 1 11.5 1v3.894a6.737 6.737 0 0 1-3 5.606a6.737 6.737 0 0 0-3 5.606V21.5H5A2.5 2.5 0 0 1 2.5 19v-4m5.5.5h2.5v.5c0 1.5 0 2.5.75 4c0 0 .75 1.5 1.75 1.5"/>`),
		g.Group(children),
	)
}
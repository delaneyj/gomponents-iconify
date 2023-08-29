package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 2v2h-2V2h2m0 16v2h-2v-2h2m-8-8v2H2v-2h2m0-4v2H2V6h2m4-4v2H6V2h2M4 2v2H2V2h2m12 0v2h-2V2h2m0 16v2h-2v-2h2M4 14v2H2v-2h2m0 4v2H2v-2h2m4 0v2H6v-2h2M20 2v18h-2V2h2Z"/>`),
		g.Group(children),
	)
}
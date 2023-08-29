package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 2v2h-2V2h2m0 16v2h-2v-2h2M8 2v2H6V2h2m8 0v2h-2V2h2m0 16v2h-2v-2h2m-8 0v2H6v-2h2M4 2v18H2V2h2m14 0h2v18h-2V2Z"/>`),
		g.Group(children),
	)
}
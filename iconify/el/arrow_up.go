package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M599.992 0L131.243 703.131h252.546V1200h432.422V703.131h252.546L599.992 0z"/>`),
		g.Group(children),
	)
}
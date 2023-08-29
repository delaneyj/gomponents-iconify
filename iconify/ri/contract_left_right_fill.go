package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContractLeftRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5 18l6-6l-6-6v12ZM19 6l-6 6l6 6V6Z"/>`),
		g.Group(children),
	)
}
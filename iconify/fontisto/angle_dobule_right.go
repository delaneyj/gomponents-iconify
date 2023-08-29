package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDobuleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.4 12L2.422 24L0 21.57L9.547 12L0 2.43L2.422 0zm10.4 0L12.822 24L10.4 21.57L19.947 12L10.4 2.43L12.822 0z"/>`),
		g.Group(children),
	)
}
package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.8 21.57L7.222 24L19.2 12L7.222 0L4.8 2.43L14.347 12z"/>`),
		g.Group(children),
	)
}
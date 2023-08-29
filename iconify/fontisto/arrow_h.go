package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.235 13.234H4.765v3.494L0 11.964l4.765-4.765v3.494h14.47V7.199L24 11.964l-4.765 4.765z"/>`),
		g.Group(children),
	)
}
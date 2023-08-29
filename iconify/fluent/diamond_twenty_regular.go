package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.586 11.414a2 2 0 0 1 0-2.828l6.002-6a2 2 0 0 1 2.828 0l6.002 6a2 2 0 0 1 0 2.828l-6.002 6a2 2 0 0 1-2.828 0l-6.002-6Zm.707-2.121a1 1 0 0 0 0 1.414l6.002 6a1 1 0 0 0 1.414 0l6.002-6a1 1 0 0 0 0-1.414l-6.002-6a1 1 0 0 0-1.414 0l-6.002 6Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.025 26.475a3.5 3.5 0 0 1 0-4.95l16.5-16.5a3.5 3.5 0 0 1 4.95 0l16.5 16.5a3.5 3.5 0 0 1 0 4.95l-16.5 16.5a3.5 3.5 0 0 1-4.95 0l-16.5-16.5Zm1.768-3.182a1 1 0 0 0 0 1.414l16.5 16.5a1 1 0 0 0 1.414 0l16.5-16.5a1 1 0 0 0 0-1.414l-16.5-16.5a1 1 0 0 0-1.414 0l-16.5 16.5Z"/>`),
		g.Group(children),
	)
}
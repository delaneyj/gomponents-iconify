package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlashForwardTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.464 1.03a.75.75 0 0 1 .507.932l-2.501 8.5a.75.75 0 0 1-1.44-.424l2.503-8.5a.75.75 0 0 1 .93-.507Z"/>`),
		g.Group(children),
	)
}
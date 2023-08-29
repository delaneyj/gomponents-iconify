package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DividerShortTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.25 4.75v14.5a.75.75 0 0 0 1.5 0V4.75a.75.75 0 0 0-1.5 0Z"/>`),
		g.Group(children),
	)
}
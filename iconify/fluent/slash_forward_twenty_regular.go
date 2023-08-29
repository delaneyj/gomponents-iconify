package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlashForwardTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.658 2.026a.5.5 0 0 1 .317.632l-5 15a.5.5 0 1 1-.95-.316l5-15a.5.5 0 0 1 .633-.316Z"/>`),
		g.Group(children),
	)
}
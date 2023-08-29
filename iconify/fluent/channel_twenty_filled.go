package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChannelTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.5 3a2.5 2.5 0 0 0-2.302 1.523A2 2 0 1 1 3 8.437V14.5A2.5 2.5 0 0 0 5.5 17h9a2.5 2.5 0 0 0 2.5-2.5v-9A2.5 2.5 0 0 0 14.5 3h-9ZM7 8.5a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5Zm.5 2.5h3a.5.5 0 0 1 0 1h-3a.5.5 0 0 1 0-1Zm-4-5.5a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}
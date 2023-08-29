package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Power(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M439.6 0H204.9L55.4 256h149.5l-128 256l341.3-320H247.5z"/>`),
		g.Group(children),
	)
}
package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m13.5 7.5l-4 4m4-4l-4-4m4 4H3M1.5 1v13"/>`),
		g.Group(children),
	)
}
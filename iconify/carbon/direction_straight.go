package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionStraight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 4l-7 7l1.414 1.414L15 7.828V28h2V7.828l4.586 4.586L23 11l-7-7z"/>`),
		g.Group(children),
	)
}
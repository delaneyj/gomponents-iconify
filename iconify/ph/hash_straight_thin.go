package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashStraightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 156h-52v-56h52a4 4 0 0 0 0-8h-52V40a4 4 0 0 0-8 0v52h-56V40a4 4 0 0 0-8 0v52H40a4 4 0 0 0 0 8h52v56H40a4 4 0 0 0 0 8h52v52a4 4 0 0 0 8 0v-52h56v52a4 4 0 0 0 8 0v-52h52a4 4 0 0 0 0-8Zm-116 0v-56h56v56Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoleculeTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 14a6 6 0 1 0-5.364-3.31l-3.359 2.363a4.5 4.5 0 1 0-.3 5.302l5.203 3.035a3.5 3.5 0 1 0 .77-1.287l-5.226-3.049a4.49 4.49 0 0 0 .156-2.591l3.592-2.526A5.986 5.986 0 0 0 19 14Z"/>`),
		g.Group(children),
	)
}
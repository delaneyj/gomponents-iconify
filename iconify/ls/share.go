package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m623 474l75-74v265H0V133h281l-84 72H76v388h547V474zm-73-121v110l238-222L550 28v109c-460 0-427 427-427 427c90-298 427-211 427-211z"/>`),
		g.Group(children),
	)
}
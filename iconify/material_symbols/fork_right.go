package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForkRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21V6.825L7.4 8.4L6 7l4-4l4 4l-1.4 1.425l-1.6-1.6v6.525q.875-.775 1.975-1.113t2.2-.337q.275 0 .525.025t.475.075L14.6 10.4L16 9l4 4l-4 4l-1.4-1.4l1.575-1.6q-.275-.05-.55-.088t-.55-.037q-1.35 0-2.488.763T11 17v4H9Z"/>`),
		g.Group(children),
	)
}
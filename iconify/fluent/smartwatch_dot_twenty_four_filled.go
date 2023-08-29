package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartwatchDotTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.5 2a2 2 0 0 1 1.79 1.105a3.487 3.487 0 0 0-.231 2.037A4.005 4.005 0 0 0 15 5H9a4 4 0 0 0-1.5.29V4a2 2 0 0 1 2-2h5Zm.5 4c.589 0 1.138.17 1.602.463c.25.368.567.686.935.935A2.99 2.99 0 0 1 18 9v1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1v2a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9a3 3 0 0 1 3-3h6ZM9 19a4 4 0 0 1-1.5-.29V20a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2v-1.29A4 4 0 0 1 15 19H9ZM22 4.5a2.5 2.5 0 1 0-5 0a2.5 2.5 0 0 0 5 0Z"/>`),
		g.Group(children),
	)
}
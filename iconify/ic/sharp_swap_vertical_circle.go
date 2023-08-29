package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSwapVerticalCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10s10-4.48 10-10S17.52 2 12 2zM6.5 9L10 5.5L13.5 9H11v4H9V9H6.5zm7.5 9.5L10.5 15H13v-4h2v4h2.5L14 18.5z"/>`),
		g.Group(children),
	)
}
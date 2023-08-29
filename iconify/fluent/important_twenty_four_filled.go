package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImportantTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2.002a3.875 3.875 0 0 0-3.875 3.875c0 2.92 1.207 6.552 1.813 8.199a2.187 2.187 0 0 0 2.064 1.423c.904 0 1.739-.542 2.063-1.418c.606-1.64 1.81-5.254 1.81-8.204A3.875 3.875 0 0 0 12 2.002ZM12.001 17a2.501 2.501 0 1 0 0 5.002a2.501 2.501 0 0 0 0-5.002Z"/>`),
		g.Group(children),
	)
}
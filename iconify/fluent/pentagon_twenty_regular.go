package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.392 1.648a2.5 2.5 0 0 1 3.37.01l5.43 4.98a2.5 2.5 0 0 1 .68 2.64l-2.027 6.02a2.5 2.5 0 0 1-2.37 1.701H6.65a2.5 2.5 0 0 1-2.351-1.65l-2.15-5.947A2.5 2.5 0 0 1 2.822 6.7l5.57-5.052Zm2.693.747a1.5 1.5 0 0 0-2.021-.006L3.492 7.44a1.5 1.5 0 0 0-.403 1.62l2.15 5.948a1.5 1.5 0 0 0 1.41.99h6.827a1.5 1.5 0 0 0 1.421-1.02l2.027-6.02a1.5 1.5 0 0 0-.408-1.583l-5.43-4.981Z"/>`),
		g.Group(children),
	)
}
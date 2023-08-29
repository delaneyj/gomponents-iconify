package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WadOfMoneyBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.25 5c-2.317.006-3.557.063-4.472.674a4 4 0 0 0-1.104 1.104C2 7.787 2 9.19 2 12c0 2.809 0 4.213.674 5.222a4 4 0 0 0 1.104 1.104c.915.611 2.155.668 4.472.674v-4.095a3.001 3.001 0 0 1 0-5.81V5Zm1.5 14h4.5V5h-4.5v14Zm6-14v4.095a3.001 3.001 0 0 1 0 5.81V19c2.317-.006 3.558-.063 4.472-.674a4.003 4.003 0 0 0 1.104-1.104C22 16.213 22 14.81 22 12c0-2.809 0-4.213-.674-5.222a4.002 4.002 0 0 0-1.104-1.104c-.915-.611-2.155-.668-4.472-.673Z"/>`),
		g.Group(children),
	)
}
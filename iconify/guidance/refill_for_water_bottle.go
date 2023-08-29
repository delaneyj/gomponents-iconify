package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefillForWaterBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 9h2v3c2.5 2 2.5 3.731 2.5 6.472V21.5a2 2 0 0 1-2 2H0M24 9h-2v3c-2.5 2-2.5 3.731-2.5 6.472V21.5a2 2 0 0 0 2 2H24M7.5 18.472V21.5a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2v-3.028c0-2.74 0-4.472-2.5-6.472V9h-4v3c-2.5 2-2.5 3.731-2.5 6.472ZM14 4.446c0 1.12-.895 2.054-2 2.054c-1.104 0-2-.934-2-2.054c0-.109.018-.222.048-.336c.19-.731.697-1.326 1.166-1.91l.12-.148C11.993 1.119 12 .342 12 0c0 .342.007 1.12.667 2.052l.12.148c.468.584.975 1.179 1.165 1.91c.03.114.048.227.048.336Z"/>`),
		g.Group(children),
	)
}
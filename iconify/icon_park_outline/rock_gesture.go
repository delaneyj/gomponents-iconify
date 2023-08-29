package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RockGesture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 25V7.5c0-1.281.5-3.5 3-3.5s3 2.219 3 3.5V30l7-7c1.297-1.297 3.078-1.922 4.5-.5c1.422 1.422 1.594 2.906 0 4.5L35 33.5S29.094 44 26 44H13s-3 0-5-2s-2-4.5-2-4.5V12.781C6 12.062 6.5 10 9 10s3 2 3 2.781V25"/><rect width="6" height="12" x="12" y="19" rx="3"/><rect width="6" height="12" x="18" y="19" rx="3"/></g>`),
		g.Group(children),
	)
}
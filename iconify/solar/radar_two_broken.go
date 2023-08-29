package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadarTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9.412 21.659c.53.142 1.06.238 1.588.292c4.795.488 9.372-2.558 10.66-7.363C23.088 9.253 19.922 3.77 14.587 2.34a9.968 9.968 0 0 0-7.501.95M12 12L5.002 6.335c-.43-.347-1.063-.283-1.366.178a9.99 9.99 0 0 0-1.295 2.898a9.997 9.997 0 0 0 2.661 9.734m2.367-10.96A6 6 0 1 0 17.811 10.5m-7.708-4.195A5.993 5.993 0 0 1 12 6c1.093 0 2.117.292 3 .802"/>`),
		g.Group(children),
	)
}
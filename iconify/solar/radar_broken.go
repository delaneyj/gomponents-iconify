package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadarBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9.412 21.659c.53.142 1.06.238 1.588.292c4.795.488 9.372-2.558 10.66-7.363C23.088 9.253 19.922 3.77 14.587 2.34a9.968 9.968 0 0 0-7.501.95M12 12L5.002 6.335c-.43-.347-1.063-.283-1.366.178a9.99 9.99 0 0 0-1.295 2.898a9.997 9.997 0 0 0 2.661 9.734"/>`),
		g.Group(children),
	)
}
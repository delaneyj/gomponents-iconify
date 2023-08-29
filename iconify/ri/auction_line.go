package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AuctionLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.005 20.003v2h-12v-2h12ZM14.59.689l7.778 7.778l-1.414 1.414l-1.061-.353l-2.475 2.475l5.657 5.657l-1.414 1.414l-5.657-5.657L13.6 15.82l.283 1.131l-1.415 1.415l-7.778-7.779l1.414-1.414l1.132.283l6.293-6.293l-.353-1.06L14.59.688Zm.707 3.535l-7.071 7.072l3.535 3.535l7.071-7.071l-3.535-3.535Z"/>`),
		g.Group(children),
	)
}
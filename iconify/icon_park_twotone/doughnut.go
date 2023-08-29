package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doughnut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDoughnut0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="24" r="19"/><circle cx="24" cy="24" r="7" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 28s3.389-2.958 6-1c4 3 6 1 6 1m11 1c.667-1 4-4.286 7-3c4 1.714 7 0 7-1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDoughnut0)"/>`),
		g.Group(children),
	)
}
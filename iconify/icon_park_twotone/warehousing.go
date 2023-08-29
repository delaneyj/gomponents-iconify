package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warehousing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWarehousing0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M34 24H18m6-6l-6 6l6 6"/><circle cx="38" cy="24" r="4" fill="#555"/><path stroke-linecap="round" d="M40.706 13A20.102 20.102 0 0 0 38 9.717A19.935 19.935 0 0 0 24 4C12.954 4 4 12.954 4 24s8.954 20 20 20c5.45 0 10.392-2.18 14-5.717A20.104 20.104 0 0 0 40.706 35"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWarehousing0)"/>`),
		g.Group(children),
	)
}
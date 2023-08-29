package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDeleteFour0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 11h32M18 5h12"/><path fill="#555" d="M12 17h24v23a3 3 0 0 1-3 3H15a3 3 0 0 1-3-3V17Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDeleteFour0)"/>`),
		g.Group(children),
	)
}
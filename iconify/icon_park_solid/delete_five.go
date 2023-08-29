package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDeleteFive0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" stroke-linecap="round" d="M8 11h32M18 5h12"/><path fill="#fff" stroke="#fff" d="M12 17h24v23a3 3 0 0 1-3 3H15a3 3 0 0 1-3-3V17Z"/><path stroke="#000" stroke-linecap="round" d="m20 25l8 8m0-8l-8 8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDeleteFive0)"/>`),
		g.Group(children),
	)
}
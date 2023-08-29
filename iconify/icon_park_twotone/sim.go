package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSim0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M8 4h24.889L40 11.273V44H8V4Z"/><path fill="#555" d="M33 26H15v10h18V26Z"/><path stroke-linecap="round" d="M15 12v6m6-6v6m6-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSim0)"/>`),
		g.Group(children),
	)
}
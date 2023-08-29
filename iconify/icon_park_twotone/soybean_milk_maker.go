package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoybeanMilkMaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSoybeanMilkMaker0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M35 10h5.064a1 1 0 0 1 .998.934l.867 13A1 1 0 0 1 40.93 25H32M7 5h28l-4 26H15l-3-21l-5-5Z"/><path fill="#555" d="M15 31h16l4 12H11l4-12Z"/><path d="M21 37h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSoybeanMilkMaker0)"/>`),
		g.Group(children),
	)
}
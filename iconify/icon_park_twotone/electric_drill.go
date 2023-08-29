package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricDrill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTElectricDrill0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M20 9h19.698a4 4 0 0 1 3.99 4.29l-.584 8a4 4 0 0 1-3.99 3.71H20V9Z"/><path d="M30.09 25H39l-4.089 11.244A4.186 4.186 0 0 1 30.977 39c-2.905 0-4.927-2.887-3.934-5.617L30.09 25Z"/><path fill="#555" d="M14 12h6v10h-6z"/><path stroke-linecap="round" d="M14 17H4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTElectricDrill0)"/>`),
		g.Group(children),
	)
}
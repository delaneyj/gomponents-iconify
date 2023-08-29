package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightHouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLightHouse0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" stroke-linecap="round" d="M6 44h36M17 20h14"/><path fill="#fff" stroke="#fff" d="M19 20h10l3 24H16l3-24Z"/><path stroke="#fff" stroke-linecap="round" d="m19 9l2 11h6l2-11"/><path stroke="#fff" stroke-linecap="round" d="m32 12l-3-3l-5-5l-5 5l-3 3m21-5l3-3M11 7L8 4m29 15l3 3m-29-3l-3 3m30-9h4m-32 0H6"/><path stroke="#000" stroke-linecap="round" d="M18 28h12m-13 8h14"/><path stroke="#fff" d="m29 20l3 24M19 20l-3 24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLightHouse0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiniSdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMiniSdCard0"><g fill="none" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-linejoin="round" d="M13.998 18.739L8 21.923V44h32V4H13.998v14.739Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M21 12v6m12-6v6m-6-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMiniSdCard0)"/>`),
		g.Group(children),
	)
}
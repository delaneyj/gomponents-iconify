package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiniSdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMiniSdCard0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" fill-rule="evenodd" stroke-linejoin="round" d="M13.998 18.739L8 21.923V44h32V4H13.998v14.739Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M21 12v6m12-6v6m-6-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMiniSdCard0)"/>`),
		g.Group(children),
	)
}
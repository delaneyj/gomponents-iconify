package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchNintendo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSwitchNintendo0"><g fill="none"><path fill="#555" stroke="#fff" stroke-width="4" d="M6 12a8 8 0 0 1 8-8h5a2 2 0 0 1 2 2v34a2 2 0 0 1-2 2h-5a8 8 0 0 1-8-8V12Zm36 2a8 8 0 0 0-8-8h-6a2 2 0 0 0-2 2v34a2 2 0 0 0 2 2h6a8 8 0 0 0 8-8V14Z"/><rect width="5" height="5" x="11" y="30" fill="#fff" rx="2.5"/><rect width="5" height="5" x="31.5" y="14" fill="#fff" rx="2.5"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15 17h-3m22 11v6m-3-3h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSwitchNintendo0)"/>`),
		g.Group(children),
	)
}
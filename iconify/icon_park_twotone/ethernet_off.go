package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EthernetOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTEthernetOff0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" stroke-linejoin="round" rx="3"/><path fill="#555" stroke-linejoin="round" d="M19 27h10v6H19zm-5-8h20v8H14z"/><path d="M33 19v-4m-6 4v-4m-6 4v-4m-6 4v-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTEthernetOff0)"/>`),
		g.Group(children),
	)
}
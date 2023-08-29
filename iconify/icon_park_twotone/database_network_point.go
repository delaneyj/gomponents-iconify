package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseNetworkPoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDatabaseNetworkPoint0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 36v-6m-4 10H6m22 0h14"/><path fill="#555" d="M28 40a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/><path d="M37 17c0 7.18-5.82 13-13 13s-13-5.82-13-13m26 0c0-7.18-5.82-13-13-13S11 9.82 11 17m26 0H11"/><path fill="#555" d="M29 17c0 7.18-2.239 13-5 13s-5-5.82-5-13s2.239-13 5-13s5 5.82 5 13Z"/><path d="M37 17H11"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDatabaseNetworkPoint0)"/>`),
		g.Group(children),
	)
}
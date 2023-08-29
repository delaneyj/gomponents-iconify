package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectAddressTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTConnectAddressTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 24c0 9.941 8.059 18 18 18s18-8.059 18-18M24 14v28M6 24h6m24 0h6"/><circle cx="24" cy="10" r="4" fill="#555"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTConnectAddressTwo0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Consignment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSConsignment0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M8 14a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2V14Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M16 12v20m16-20v20"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M28 12h8m-24 0h8m-8 20h8m8 0h8M4 38h40"/><path stroke="#fff" stroke-linecap="round" d="M18 38v2m-6-2v2m-6-2v2m18-2v2m6-2v2m6-2v2m6-2v2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M30 12V6H18v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSConsignment0)"/>`),
		g.Group(children),
	)
}
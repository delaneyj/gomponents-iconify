package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTListOne0"><g fill="#555" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M18 6h16v6H18V6Zm0 15h20v6H18v-6Zm0 15h26v6H18v-6Z"/><circle cx="8" cy="9" r="2"/><circle cx="8" cy="24" r="2"/><circle cx="8" cy="39" r="2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTListOne0)"/>`),
		g.Group(children),
	)
}
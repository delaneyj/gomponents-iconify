package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMessageSearch0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M25.5 37H21l-10 5v-5H4V7h40v11"/><circle cx="34" cy="28" r="6" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="m39 32l5 4M12 15h6m-6 6h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMessageSearch0)"/>`),
		g.Group(children),
	)
}
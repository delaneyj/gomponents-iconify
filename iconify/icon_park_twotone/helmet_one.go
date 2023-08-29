package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelmetOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHelmetOne0"><g fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 11c-9.389 0-17 7.815-17 17.454V35h34v-6.546C41 18.814 33.389 11 24 11ZM4 35h40v6H4z"/><path d="M21 6h6v18h-6z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHelmetOne0)"/>`),
		g.Group(children),
	)
}
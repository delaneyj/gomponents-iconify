package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Castle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCastle0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m4 11l7-7l7 7H4Zm26 0l7-7l7 7H30Z"/><path fill="#555" d="M44 44V26h-4v-9h-6v9h-6v-5l-4-4l-4 4v5h-6v-9H8v9H4v18h14v-4a6 6 0 0 1 12 0v4h14Z"/><path d="M7 11h8v6H7zm26 0h8v6h-8z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCastle0)"/>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Badge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBadge0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M38 22v18H8V10h18"/><path fill="#555" d="M38 14a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBadge0)"/>`),
		g.Group(children),
	)
}
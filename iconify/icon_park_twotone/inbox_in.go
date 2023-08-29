package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTInboxIn0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="m5 30l5-24h28l5 24"/><path fill="#555" d="M5 30h9.91l1.817 6h14.546l1.818-6H43v13H5V30Z"/><path stroke-linecap="round" d="m18 20l6 6l6-6m-6 6V14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTInboxIn0)"/>`),
		g.Group(children),
	)
}
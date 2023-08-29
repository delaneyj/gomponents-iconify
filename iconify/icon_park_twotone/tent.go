package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTent0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 12L4 36h12"/><path fill="#555" d="M38 12H10l6 24h28l-6-24Z"/><path d="M12 18h27m-29-6l3 12m25-12l3 12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTent0)"/>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRadio0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="38" height="28" x="5" y="13" fill="#555" rx="2"/><circle cx="18" cy="28" r="6" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="M30 24h6m-6 8h6M7 13l18-9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRadio0)"/>`),
		g.Group(children),
	)
}
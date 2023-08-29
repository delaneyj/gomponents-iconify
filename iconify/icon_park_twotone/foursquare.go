package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foursquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFoursquare0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M38 4H14v40l10-18h8l6-22Z"/><path d="M35 15H25m11.636-6l-3.273 12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFoursquare0)"/>`),
		g.Group(children),
	)
}
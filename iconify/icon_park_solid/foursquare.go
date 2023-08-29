package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foursquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFoursquare0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M38 4H14v40l10-18h8l6-22Z"/><path stroke="#000" d="M35 15H25"/><path stroke="#fff" d="m36.636 9l-3.273 12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFoursquare0)"/>`),
		g.Group(children),
	)
}
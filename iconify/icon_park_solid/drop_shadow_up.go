package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropShadowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDropShadowUp0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#fff" fill-rule="evenodd" d="M24 43a15.95 15.95 0 0 0 11.314-4.686A15.95 15.95 0 0 0 40 27a15.95 15.95 0 0 0-4.686-11.314A15.95 15.95 0 0 0 24 11a15.95 15.95 0 0 0-11.314 4.686A15.95 15.95 0 0 0 8 27a15.95 15.95 0 0 0 4.686 11.314A15.95 15.95 0 0 0 24 43Z" clip-rule="evenodd"/><path d="M39.557 11.444A21.931 21.931 0 0 0 24 5a21.931 21.931 0 0 0-15.556 6.444"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDropShadowUp0)"/>`),
		g.Group(children),
	)
}
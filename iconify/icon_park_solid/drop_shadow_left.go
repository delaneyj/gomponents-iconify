package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropShadowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDropShadowLeft0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#fff" fill-rule="evenodd" d="M27 40a15.95 15.95 0 0 0 11.314-4.686A15.95 15.95 0 0 0 43 24a15.95 15.95 0 0 0-4.686-11.314A15.95 15.95 0 0 0 27 8a15.95 15.95 0 0 0-11.314 4.686A15.95 15.95 0 0 0 11 24a15.95 15.95 0 0 0 4.686 11.314A15.95 15.95 0 0 0 27 40Z" clip-rule="evenodd"/><path d="M11.444 8.444A21.931 21.931 0 0 0 5 24a21.931 21.931 0 0 0 6.444 15.556"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDropShadowLeft0)"/>`),
		g.Group(children),
	)
}
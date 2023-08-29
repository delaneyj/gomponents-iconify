package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropShadowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDropShadowRight0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#fff" fill-rule="evenodd" d="M21 40a15.95 15.95 0 0 0 11.314-4.686A15.95 15.95 0 0 0 37 24a15.95 15.95 0 0 0-4.686-11.314A15.95 15.95 0 0 0 21 8a15.95 15.95 0 0 0-11.314 4.686A15.95 15.95 0 0 0 5 24a15.95 15.95 0 0 0 4.686 11.314A15.95 15.95 0 0 0 21 40Z" clip-rule="evenodd"/><path d="M36.557 39.556A21.932 21.932 0 0 0 43 24a21.932 21.932 0 0 0-6.443-15.556"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDropShadowRight0)"/>`),
		g.Group(children),
	)
}
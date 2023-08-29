package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropShadowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDropShadowDown0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" fill-rule="evenodd" d="M24 37a15.95 15.95 0 0 0 11.314-4.686A15.95 15.95 0 0 0 40 21a15.95 15.95 0 0 0-4.686-11.314A15.95 15.95 0 0 0 24 5a15.95 15.95 0 0 0-11.314 4.686A15.95 15.95 0 0 0 8 21a15.95 15.95 0 0 0 4.686 11.314A15.95 15.95 0 0 0 24 37Z" clip-rule="evenodd"/><path d="M39.556 36.556A21.932 21.932 0 0 1 24 43a21.932 21.932 0 0 1-15.557-6.444"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDropShadowDown0)"/>`),
		g.Group(children),
	)
}
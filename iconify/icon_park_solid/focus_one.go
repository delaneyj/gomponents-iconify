package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FocusOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFocusOne0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 6H8a2 2 0 0 0-2 2v8m10 26H8a2 2 0 0 1-2-2v-8m26 10h8a2 2 0 0 0 2-2v-8M32 6h8a2 2 0 0 1 2 2v8"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="4" d="M24 31a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="4" d="M24 17v-4m0 22v-4m11-7h-4m-14 0h-4"/><path fill="#000" d="M24 26a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFocusOne0)"/>`),
		g.Group(children),
	)
}
package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Focus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFocus0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 6H8a2 2 0 0 0-2 2v8m10 26H8a2 2 0 0 1-2-2v-8m26 10h8a2 2 0 0 0 2-2v-8M32 6h8a2 2 0 0 1 2 2v8"/><rect width="20" height="20" x="14" y="14" fill="#fff" stroke="#fff" stroke-width="4" rx="10"/><circle r="3" fill="#000" transform="matrix(-1 0 0 1 24 24)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFocus0)"/>`),
		g.Group(children),
	)
}
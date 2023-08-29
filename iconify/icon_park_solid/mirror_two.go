package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirrorTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMirrorTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="17" r="13" fill="#fff"/><path d="M42 17c0 9.941-8.059 18-18 18S6 26.941 6 17m36 0h-4m-28 0H6m24 27H18m6 0v-8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMirrorTwo0)"/>`),
		g.Group(children),
	)
}
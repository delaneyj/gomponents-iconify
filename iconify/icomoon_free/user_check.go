package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 9.5L10.5 14L9 12.5l-1 1l2.5 2.5l5.5-5.5z"/><path fill="currentColor" d="M7 12h5v-1.799c-1.05-.613-2.442-1.033-4-1.16v-.825c1.102-.621 2-2.168 2-3.716C10 2.015 10 0 7 0S4 2.015 4 4.5c0 1.548.898 3.095 2 3.716v.825C2.608 9.318 0 10.985 0 13h7v-1z"/>`),
		g.Group(children),
	)
}
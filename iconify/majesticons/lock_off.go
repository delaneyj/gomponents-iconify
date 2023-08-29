package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M7 7.828V9H6a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h12c.872 0 1.657-.372 2.205-.966l-7.223-7.223A.97.97 0 0 1 13 14v3a1 1 0 1 1-2 0v-3a1 1 0 0 1 1.19-.982L7 7.828zm14 8.344V12a3 3 0 0 0-3-3h-1V7a5 5 0 0 0-8.62-3.449l1.415 1.415A3 3 0 0 1 15 7v2h-1.172L21 16.172z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18"/></g>`),
		g.Group(children),
	)
}
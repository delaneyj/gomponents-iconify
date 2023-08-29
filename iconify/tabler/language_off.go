package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 5h1m4 0h2M9 3v2m-.508 3.517C7.678 11.172 5.972 13 4 13"/><path d="M5 9c0 2.144 2.952 3.908 6.7 4m.3 7l2.463-5.541m1.228-2.764L16 11l.8 1.8M18 18h-5.1M3 3l18 18"/></g>`),
		g.Group(children),
	)
}
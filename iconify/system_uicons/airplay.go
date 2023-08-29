package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m6.5 14.5l-1-.035c-1.102-.003-2-.932-2-2.034V6.465a2 2 0 0 1 2-2l10-.002a2 2 0 0 1 2 2v6.002a2 2 0 0 1-2 2l-1 .037"/><path d="m10.5 13.5l-3 3h6z"/></g>`),
		g.Group(children),
	)
}
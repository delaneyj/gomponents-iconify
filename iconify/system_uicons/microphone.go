package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m10.39 2.615l.11-.004a2.893 2.893 0 0 1 3 2.891V9.5a3 3 0 1 1-6 0V5.613a3 3 0 0 1 2.89-2.998z"/><path d="M15.5 9.5a5 5 0 0 1-9.995.217L5.5 9.5m5 5v4"/></g>`),
		g.Group(children),
	)
}
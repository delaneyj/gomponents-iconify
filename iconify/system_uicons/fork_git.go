package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForkGit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m14.5 3.5l3 3l-3 3"/><path d="M17.5 6.5h-5l-4 5.086m6 .914l3 3l-3 3"/><path d="M17.5 15.5h-5l-4-4h-6"/></g>`),
		g.Group(children),
	)
}
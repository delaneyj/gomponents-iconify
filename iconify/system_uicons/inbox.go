package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.16 4.5h8.68a1 1 0 0 1 .92.606L18.5 11.5v4a2 2 0 0 1-2 2h-12a2 2 0 0 1-2-2v-4l2.74-6.394a1 1 0 0 1 .92-.606z"/><path d="M2.5 11.5h4a1 1 0 0 1 1 1v1a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1h4"/></g>`),
		g.Group(children),
	)
}
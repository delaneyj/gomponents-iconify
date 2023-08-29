package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 16a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/><path d="M12 21H7a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2M8 11V7a4 4 0 1 1 8 0v4m0 11l5-5m0 4.5V17h-4.5"/></g>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 3v4a1 1 0 0 0 1 1h4"/><path d="M12 21H7a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7l5 5v4.5"/><path d="M14 17.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 1 0-5 0m4.5 2L21 22"/></g>`),
		g.Group(children),
	)
}
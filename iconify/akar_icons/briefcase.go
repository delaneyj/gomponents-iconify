package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="3" y="7" rx="2"/><path d="M9 6a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v1H9V6Zm1 6l.211.106a4 4 0 0 0 3.578 0L14 12"/></g>`),
		g.Group(children),
	)
}
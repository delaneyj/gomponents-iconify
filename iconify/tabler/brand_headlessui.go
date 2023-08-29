package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandHeadlessui(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m6.744 4.325l7.82-1.267a4.456 4.456 0 0 1 5.111 3.686l1.267 7.82a4.456 4.456 0 0 1-3.686 5.111l-7.82 1.267a4.456 4.456 0 0 1-5.111-3.686l-1.267-7.82a4.456 4.456 0 0 1 3.686-5.111z"/><path d="m7.252 7.704l7.897-1.28a1 1 0 0 1 1.147.828l.36 2.223l-9.562 3.51l-.67-4.134a1 1 0 0 1 .828-1.147z"/></g>`),
		g.Group(children),
	)
}
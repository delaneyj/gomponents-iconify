package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingCorrectSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.974 19.375a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25Zm4.226 3.75a4.473 4.473 0 0 0-8.449 0m18.275-3.75a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25Zm4.224 3.75a4.473 4.473 0 0 0-8.449 0M.75 6.875h3.5m-1.5 2l-2-2l2-2m20.5 2h-3.5m1.5 2l2-2l-2-2m-9.25 6a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path d="m10 6.375l1.083 1.083a.5.5 0 0 0 .76-.063L14 4.375"/></g>`),
		g.Group(children),
	)
}
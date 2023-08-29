package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M4.974 17.375a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25Zm4.226 3.75a4.473 4.473 0 0 0-8.449 0m18.275-3.75a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25Zm4.224 3.75a4.473 4.473 0 0 0-8.449 0"/><path d="M6 3.75A.375.375 0 1 1 6 3m0 .75A.375.375 0 1 0 6 3m6 .75A.375.375 0 0 1 12 3m0 .75A.375.375 0 0 0 12 3m6 .75A.375.375 0 0 1 18 3m0 .75A.375.375 0 0 0 18 3"/></g>`),
		g.Group(children),
	)
}
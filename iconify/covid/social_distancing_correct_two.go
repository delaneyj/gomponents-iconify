package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingCorrectTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.974 19.375a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25Zm4.226 3.75a4.473 4.473 0 0 0-8.449 0m18.275-3.75a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25Zm4.224 3.75a4.473 4.473 0 0 0-8.449 0M5 8.875v2m14-2v2m-14-1h14m-10-7l2 2l4-4"/>`),
		g.Group(children),
	)
}
package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingCorrectOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.5 19.375a3.375 3.375 0 1 0 0-6.75a3.375 3.375 0 0 0 0 6.75Zm4.75 3.75a6.027 6.027 0 0 0-9.5 0"/><path d="M21.875 16.111a6.762 6.762 0 0 1-6.443-1.511M5.5 19.375a3.375 3.375 0 1 0 0-6.75a3.375 3.375 0 0 0 0 6.75Zm4.749 3.75a6.026 6.026 0 0 0-9.5 0"/><path d="M8.875 16.111A6.762 6.762 0 0 1 2.432 14.6M5 7.875v2m14-2v2m-14-1h14m-10-6l2 2l4-4"/></g>`),
		g.Group(children),
	)
}
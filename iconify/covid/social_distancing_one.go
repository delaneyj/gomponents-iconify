package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M18.5 17.375a3.375 3.375 0 1 0 0-6.75a3.375 3.375 0 0 0 0 6.75Zm4.75 3.75a6.027 6.027 0 0 0-9.5 0"/><path stroke-linecap="round" stroke-linejoin="round" d="M21.875 14.111a6.762 6.762 0 0 1-6.443-1.511M5.5 17.375a3.375 3.375 0 1 0 0-6.75a3.375 3.375 0 0 0 0 6.75Zm4.749 3.75a6.026 6.026 0 0 0-9.5 0"/><path stroke-linecap="round" stroke-linejoin="round" d="M8.875 14.111A6.762 6.762 0 0 1 2.432 12.6"/><path d="M5.875 3.75a.375.375 0 1 1 0-.75m0 .75a.375.375 0 1 0 0-.75M12 3.75A.375.375 0 0 1 12 3m0 .75A.375.375 0 0 0 12 3m6.125.75a.375.375 0 0 1 0-.75m0 .75a.375.375 0 0 0 0-.75"/></g>`),
		g.Group(children),
	)
}
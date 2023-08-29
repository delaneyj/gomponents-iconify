package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingDoNotCloseThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.218 10.25a4.75 4.75 0 1 0 0-9.5a4.75 4.75 0 0 0 0 9.5ZM17.021.946a4.75 4.75 0 1 1 .642 9.251M6.471 17.724a2.125 2.125 0 0 1-3 0l-2.1-2.1a2.125 2.125 0 0 1 3.005-3l2.1 2.1a2.125 2.125 0 0 1-.005 3Zm-3.498 3.034v2.492"/><path d="M4.251 12.508A6.235 6.235 0 0 1 9.06 10.25h3.222a6.3 6.3 0 0 1 2.815.656m2.913 12.319a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm3.535-8.535l-7.07 7.07"/></g>`),
		g.Group(children),
	)
}
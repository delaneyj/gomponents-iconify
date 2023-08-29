package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialDistancingNotAllowedSpaceWoman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m1 19.05l4 4m0-4l-4 4m18-4l4 4m0-4l-4 4m-7-16.3a3 3 0 1 0 0-6a3 3 0 0 0 0 6ZM9.114 9.114A5.246 5.246 0 0 0 6.75 13.5v2.25H9l.75 7.5h4.5l.75-7.5h2.25V13.5a5.246 5.246 0 0 0-2.364-4.386L12 13.5L9.114 9.114Z"/><path d="M9 3.75S9 7.5 6.75 7.5M15 3.75s0 3.75 2.25 3.75"/></g>`),
		g.Group(children),
	)
}
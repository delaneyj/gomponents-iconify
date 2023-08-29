package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayVolleyball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPlayVolleyball0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-miterlimit="2" stroke-width="4" d="M29 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path fill="#fff" d="M41 24a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4" d="m16 24l-2.27 11.26a2 2 0 0 1-2.81 1.5L4 34m34-4L23.76 19.07c-.77-.57-1.83-.51-2.54.13l-4.84 4.43a.996.996 0 0 0 .09 1.55L26 32.15c.72.53 1.01 1.48.69 2.32L23.12 44"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPlayVolleyball0)"/>`),
		g.Group(children),
	)
}
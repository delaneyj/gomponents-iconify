package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChairLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M5.032 6.938c-.275-2.21 1.27-4.24 3.51-4.616l.398-.067c2.025-.34 4.095-.34 6.12 0l.398.067c2.24.376 3.785 2.407 3.51 4.616l-.017.145A1.058 1.058 0 0 1 17.893 8H6.107c-.539 0-.992-.394-1.057-.917l-.018-.145Z"/><path stroke-linecap="round" d="M9 8L6 22m9-14l3 14"/><path stroke-linecap="round" d="M16.5 17h-9" opacity=".5"/></g>`),
		g.Group(children),
	)
}
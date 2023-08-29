package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Waist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.4 4s-1.6 3.75-1.6 6.857c0 .995.34 1.827.8 2.656c.528.954 1.214 1.903 1.717 3.09A8.49 8.49 0 0 1 20 20M5.6 4s1.6 3.75 1.6 6.857c0 .995-.34 1.827-.8 2.656c-.528.954-1.214 1.903-1.717 3.09A8.483 8.483 0 0 0 4 20m2.4-6.487h11.2"/><path d="M4.683 16.604S10.4 17.713 12 20c1.6-2.286 7.317-3.396 7.317-3.396"/></g>`),
		g.Group(children),
	)
}
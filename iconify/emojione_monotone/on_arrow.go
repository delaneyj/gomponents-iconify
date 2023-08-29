package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OnArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M18 24h27.999v10L62 18L45.999 2v10H18V2L2 18l16 16zm2 16c-3.308 0-6 2.848-6 6.346v9.307c0 3.5 2.692 6.348 6 6.348s6-2.848 6-6.348v-9.307C26 42.848 23.308 40 20 40m3 15.652c0 1.633-1.346 2.965-3 2.965s-3-1.332-3-2.965v-9.307c0-1.633 1.346-2.963 3-2.963s3 1.33 3 2.963v9.307m15.999-1.357L33.152 40H30v22h3V47.703L38.847 62h3.152V40h-3zm7 3.705h4v4h-4zm.736-4h2.528l.736-14h-4z"/>`),
		g.Group(children),
	)
}
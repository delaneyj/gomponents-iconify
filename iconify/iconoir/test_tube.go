package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.141 19.995c2.458 1.72 4.281-.012 5.318-1.492l7.3-10.426l1.966-1.065l-6.553-4.588l-8.447 12.064c-1.037 1.48-2.041 3.786.416 5.507Z"/><path d="M16.091 11.02c-2.876-.853-4.403.781-7.28-.071"/></g>`),
		g.Group(children),
	)
}
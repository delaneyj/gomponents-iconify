package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntersectAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15.01 3l-.01.011M11.01 3l-.01.011M7.01 3L7 3.011M3.01 3L3 3.011M3.01 7L3 7.011M3.01 11l-.01.011M3.01 15l-.01.011m6 5.999l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011m3.99.011l.01-.011M21 17.01l.01-.011M21 13.01l.01-.011M21 9.01l.01-.011M9 17v-7a1 1 0 0 1 1-1h7"/><path d="M15 7v7a1 1 0 0 1-1 1H7"/></g>`),
		g.Group(children),
	)
}
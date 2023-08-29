package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderInner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 20.01l.01-.011M4 16.01l.01-.011M4 8.01l.01-.011M4 4.01l.01-.011M8 4.01l.01-.011M16 4.01l.01-.011M20 4.01l.01-.011M20 8.01l.01-.011M8 20.01l.01-.011m7.99.011l.01-.011m3.99.011l.01-.011M20 16.01l.01-.011M4 12h8m8 0h-8m0 0V4m0 8v8"/>`),
		g.Group(children),
	)
}
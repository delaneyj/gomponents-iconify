package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18M10.012 6.016l1.981-4.014l3.086 6.253l6.9 1l-4.421 4.304m.012 4.01l.588 3.426L12 17.75l-6.172 3.245l1.179-6.873l-5-4.867l6.327-.917"/>`),
		g.Group(children),
	)
}
package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path d="M216.008 66H174V56a22.025 22.025 0 0 0-22-22h-48a22.025 22.025 0 0 0-22 22v10H40.008a14.016 14.016 0 0 0-14 14v128a14.016 14.016 0 0 0 14 14h176a14.016 14.016 0 0 0 14-14V80a14.016 14.016 0 0 0-14-14zM94 56a10.011 10.011 0 0 1 10-10h48a10.011 10.011 0 0 1 10 10v10H94zm124.008 152a2.003 2.003 0 0 1-2 2h-176a2.003 2.003 0 0 1-2-2V80a2.003 2.003 0 0 1 2-2h176a2.003 2.003 0 0 1 2 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}
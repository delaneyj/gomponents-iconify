package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cookie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M216 3Q128 3 65.5 65.5T3 216t62.5 150.5T216 429t150.5-62.5T429 216T366.5 65.5T216 3zm0 384q-70 0-120.5-50.5T45 216T95.5 95.5T216 45t120.5 50.5T387 216t-50.5 120.5T216 387zm21-299h-21q-21 0-21 21q0 22 21 22h21q22 0 22-22q0-9-6-15t-16-6zm-91 92q-16-16-30 0q-16 14 0 30l15 15q14 14 30 0q14-16 0-30zm85 91q-15-13-30 0q-14 16 0 30l15 15q15 15 30 0q14-14 0-30zm70-91q-14-16-30 0q-14 14 0 30l15 15q16 14 30 0q16-16 0-30z"/>`),
		g.Group(children),
	)
}
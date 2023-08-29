package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1190 581l293-293l-107-107l-293 293zm447-293q0 27-18 45L333 1619q-18 18-45 18t-45-18L45 1421q-18-18-18-45t18-45L1331 45q18-18 45-18t45 18l198 198q18 18 18 45zM286 98l98 30l-98 30l-30 98l-30-98l-98-30l98-30l30-98zm350 162l196 60l-196 60l-60 196l-60-196l-196-60l196-60l60-196zm930 478l98 30l-98 30l-30 98l-30-98l-98-30l98-30l30-98zM926 98l98 30l-98 30l-30 98l-30-98l-98-30l98-30l30-98z"/>`),
		g.Group(children),
	)
}
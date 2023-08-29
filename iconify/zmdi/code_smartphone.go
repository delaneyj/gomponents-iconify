package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeSmartphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M64 91v42H21V48q0-18 12.5-30.5T64 5l213 1q18 0 30.5 12T320 48v85h-43V91H64zm179 247l-30-30l68-68l-68-68l30-30l98 98zm-115-30l-30 30l-98-98l98-98l30 30l-68 68zm149 81v-42h43v85q0 18-12.5 30.5T277 475H64q-18 0-30.5-12.5T21 432v-85h43v42h213z"/>`),
		g.Group(children),
	)
}
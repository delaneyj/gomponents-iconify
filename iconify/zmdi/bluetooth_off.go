package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M192 84v69l-43-43V3h22l121 121l-64 65l-30-30l34-35zM30 45l311 312l-30 30l-49-49l-91 91h-22V267l-98 98l-30-30l120-119L0 75zm162 303l40-40l-40-41v81z"/>`),
		g.Group(children),
	)
}
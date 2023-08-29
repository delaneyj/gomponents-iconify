package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M128 512v-43h43v43h-43zm-85 0v-43h42v43H43zm170 0v-43h43v43h-43zm58-390l-92 91l92 92l-122 122h-21V265l-98 98l-30-30l119-120L0 94l30-30l98 98V0h21zM171 82v80l40-40zm40 223l-40-40v80z"/>`),
		g.Group(children),
	)
}
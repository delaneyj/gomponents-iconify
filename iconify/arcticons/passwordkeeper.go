package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Passwordkeeper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 8.28l1.65 1.79A14 14 0 0 1 24 38h0a14 14 0 0 1-1.64-28L24 8.28Zm0 9.54A6.18 6.18 0 0 1 30.18 24h0A6.18 6.18 0 0 1 24 30.18h0A6.18 6.18 0 0 1 17.82 24h0A6.18 6.18 0 0 1 24 17.82Zm0-11.64V3.5m12.6 7.9l1.9-1.9M41.82 24h2.68m-7.9 12.6l1.9 1.9M24 41.82v2.68m-12.6-7.9l-1.9 1.9M6.18 24H3.5m7.9-12.6L9.5 9.5"/>`),
		g.Group(children),
	)
}
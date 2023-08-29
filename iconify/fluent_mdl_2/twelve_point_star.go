package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwelvePointStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1754 1220l142 308l-338 30l-30 338l-308-142l-196 277l-196-277l-308 142l-30-338l-338-30l142-308l-277-196l277-196l-142-308l338-30l30-338l308 142l196-277l196 277l308-142l30 338l338 30l-142 308l277 196l-277 196zm-50 196l-111-240l216-152l-216-152l111-240l-264-24l-24-264l-240 111l-152-216l-152 216l-240-111l-24 264l-264 24l111 240l-216 152l216 152l-111 240l264 24l24 264l240-111l152 216l152-216l240 111l24-264l264-24z"/>`),
		g.Group(children),
	)
}
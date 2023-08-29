package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileReport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 384v256h-384V384h384zM0 1024h384v640H0v-640zm512-640h384v1280H512V384zm1280 384q27 0 50 10t40 27t28 41t10 50v1024q0 27-10 50t-27 40t-41 28t-50 10h-640q-27 0-50-10t-40-27t-28-41t-10-50V896q0-27 10-50t27-40t41-28t50-10h640zm0 128h-640v1024h640V896zm-512 768h384v128h-384v-128zm128-1024h-384V0h384v640z"/>`),
		g.Group(children),
	)
}
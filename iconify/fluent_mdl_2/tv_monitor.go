package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TVMonitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1408h-896v128h256v128H640v-128h256v-128H0V256h1920v1152zM128 384v896h1664V384H128z"/>`),
		g.Group(children),
	)
}
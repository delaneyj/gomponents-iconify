package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Album(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1856 256q40 0 75 15t61 41t41 61t15 75v1152q0 40-15 75t-41 61t-61 41t-75 15H0V256h1856zM128 1664h128V384H128v1280zM1920 448q0-26-19-45t-45-19H384v1280h1472q26 0 45-19t19-45V448zM768 640h768v384H768V640zm128 256h512V768H896v128z"/>`),
		g.Group(children),
	)
}
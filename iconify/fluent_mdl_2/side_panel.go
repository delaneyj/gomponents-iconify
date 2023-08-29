package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SidePanel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 256h2048v1408H0V256zm128 384v896h384V640H128zm1792 896V640H640v896h1280zM128 512h1792V384H128v128z"/>`),
		g.Group(children),
	)
}
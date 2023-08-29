package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FitWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 256h2048v1536H0V256zm1920 1408V384H128v1280h1792zM555 1109l-86 86l-236-235l236-235l86 86l-85 85h426v128H470l85 85zm938 0l85-85h-426V896h426l-85-85l86-86l236 235l-236 235l-86-86z"/>`),
		g.Group(children),
	)
}
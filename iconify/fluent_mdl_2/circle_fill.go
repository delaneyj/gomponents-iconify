package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M32 1024q0-137 35-264t100-237t155-200t201-155T760 68t264-36q137 0 264 35t237 100t200 155t155 201t100 237t36 264q0 137-35 264t-100 237t-155 200t-201 155t-237 100t-264 36q-137 0-264-35t-237-100t-200-155t-155-201t-100-237t-36-264z"/>`),
		g.Group(children),
	)
}
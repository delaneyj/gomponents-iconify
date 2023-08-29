package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1024H890q-12 56-41 103t-70 81t-94 53t-109 19q-57 0-109-19t-93-53t-71-81t-41-103H0V896h262q12-56 41-103t70-81t94-53t109-19q57 0 109 19t93 53t71 81t41 103h1158v128zM576 1152q40 0 75-15t61-41t41-61t15-75q0-40-15-75t-41-61t-61-41t-75-15q-40 0-75 15t-61 41t-41 61t-15 75q0 40 15 75t41 61t61 41t75 15z"/>`),
		g.Group(children),
	)
}
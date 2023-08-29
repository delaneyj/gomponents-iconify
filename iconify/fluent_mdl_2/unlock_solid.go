package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 896v1152H256V896h1152V512q0-82-29-152t-80-122t-122-81t-153-29q-82 0-152 29t-122 80t-81 122t-29 153H512q0-109 39-202t108-163T821 39t203-39q109 0 202 39t163 108t108 162t39 203v384h256z"/>`),
		g.Group(children),
	)
}
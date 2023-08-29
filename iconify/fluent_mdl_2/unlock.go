package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 896v1152H256V896h1152V512q0-82-29-152t-81-121t-122-81t-152-30q-82 0-152 29t-121 81t-81 122t-30 152H512q0-109 39-202t108-163T821 39t203-39q109 0 202 39t163 108t108 162t39 203v384h256zm-128 128H384v896h1280v-896z"/>`),
		g.Group(children),
	)
}
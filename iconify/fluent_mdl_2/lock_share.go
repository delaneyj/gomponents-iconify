package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1629 1789l-90-90l162-163h-293q-27 0-50 10t-40 27t-28 41t-10 50v128h-128v-128q0-53 20-99t55-82t81-55t100-20h293l-162-163l90-90l317 317l-317 317zm-477 259v-128h640v128h-640zM384 1024v896h640v128H256V896h256V522q0-108 39-203t108-166T821 41t203-41q109 0 202 41t163 112t108 166t39 203v374h256v128H384zm256-502v374h768V522q0-81-29-152t-80-126t-122-85t-153-31q-82 0-152 31t-122 85t-81 125t-29 153z"/>`),
		g.Group(children),
	)
}
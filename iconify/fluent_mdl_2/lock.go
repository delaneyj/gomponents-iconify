package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 896v1152H256V896h256V522q0-108 39-203t108-166T821 41t203-41q109 0 202 41t163 112t108 166t39 203v374h256zm-1152 0h768V522q0-81-29-152t-80-126t-122-85t-153-31q-82 0-152 31t-122 85t-81 125t-29 153v374zm1024 128H384v896h1280v-896z"/>`),
		g.Group(children),
	)
}
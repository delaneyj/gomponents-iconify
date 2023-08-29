package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 43q18 0 30.5 12.5T427 85v214q0 17-12.5 29.5T384 341H43q-18 0-30.5-12.5T0 299V85q0-17 12.5-29.5T43 43h341zm-192 64v42h43v-42h-43zm0 64v42h43v-42h-43zm-64-64v42h43v-42h-43zm0 64v42h43v-42h-43zm-21 42v-42H64v42h43zm0-64v-42H64v42h43zm192 150v-43H128v43h171zm0-86v-42h-43v42h43zm0-64v-42h-43v42h43zm64 64v-42h-43v42h43zm0-64v-42h-43v42h43z"/>`),
		g.Group(children),
	)
}
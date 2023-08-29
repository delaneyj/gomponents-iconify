package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 363v-86h427v86H0zm43-64v42h42v-42H43zM0 21h427v86H0V21zm85 64V43H43v42h42zM0 235v-86h427v86H0zm43-64v42h42v-42H43z"/>`),
		g.Group(children),
	)
}
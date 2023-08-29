package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 149v43H0v-43h256zm0-85v43H0V64h256zm85 171h86v42h-86v86h-42v-86h-86v-42h86v-86h42v86zM0 277v-42h171v42H0z"/>`),
		g.Group(children),
	)
}
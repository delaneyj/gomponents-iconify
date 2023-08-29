package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 496v-43h42v43H43zm106.5-235q-26.5 0-45.5-18.5T85 197V69q0-26 19-45t45.5-19t45 19T213 69v128q0 27-18.5 45.5t-45 18.5zM128 496v-43h43v43h-43zm85 0v-43h43v43h-43zm86-299q0 54-37.5 95T171 341v70h-43v-70q-53-8-90.5-49T0 197h36q0 47 34 78t79.5 31t79-31t33.5-78h37z"/>`),
		g.Group(children),
	)
}
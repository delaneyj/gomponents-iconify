package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149.5 259q-26.5 0-45.5-19t-19-45V67q0-27 19-45.5T149.5 3t45 18.5T213 67v128q0 26-18.5 45t-45 19zM262 195h37q0 54-37.5 94.5T171 338v70h-43v-70q-53-8-90.5-49T0 195h36q0 46 34 77t79.5 31t79-31t33.5-77z"/>`),
		g.Group(children),
	)
}
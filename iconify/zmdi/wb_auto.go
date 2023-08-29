package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WbAuto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m146 206l25-78l24 78h-49zM469 85h39l-44 192h-37l-32-130l-32 130h-38l-2-9q-21 43-62 69t-90 26q-71 0-121-50T0 192T50 71t121-50q81 0 133 64h16l26 135l32-135h34l32 135zM220 277h40L192 85h-43L81 277h41l15-42h68z"/>`),
		g.Group(children),
	)
}
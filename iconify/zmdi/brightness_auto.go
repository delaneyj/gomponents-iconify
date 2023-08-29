package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessAuto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m217 262l24-78l25 78h-49zm195-85l71 71l-71 71v100H312l-71 70l-70-70H71V319L0 248l71-71V77h100l70-70l71 70h100v100zM290 333h41l-68-192h-43l-68 192h40l15-42h68z"/>`),
		g.Group(children),
	)
}
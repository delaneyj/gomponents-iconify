package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterFrames(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 85q18 0 30.5 12.5T427 128v299q0 17-12.5 29.5T384 469H43q-18 0-30.5-12.5T0 427V128q0-18 12.5-30.5T43 85h85l85-85l86 85h85zm0 342V128h-96l-74-75l-75 75H43v299h341zm-43-256H85v213h256V171z"/>`),
		g.Group(children),
	)
}
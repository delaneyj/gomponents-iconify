package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 104c-83.813 0-152 68.187-152 152s68.187 152 152 152s152-68.187 152-152s-68.187-152-152-152Zm0 272a120 120 0 1 1 120-120a120.136 120.136 0 0 1-120 120ZM240 16h32v48h-32zm0 432h32v48h-32zm208-208h48v32h-48zm-432 0h48v32H16zm372.687 171.314l22.627-22.627l32 32l-22.627 22.627zm-320-320l22.628-22.628l32 32l-22.628 22.628zm-.002 329.375l32-32l22.628 22.626l-32 32zm320.002-320.003l32-32l22.628 22.628l-32 32z"/>`),
		g.Group(children),
	)
}
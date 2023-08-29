package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M144 144H16v32h160V16h-32v128zm224 0V16h-32v160h160v-32H368zm-32 352h32V368h128v-32H336v160zM16 368h128v128h32V336H16v32z"/>`),
		g.Group(children),
	)
}
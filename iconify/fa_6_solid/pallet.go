package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 320c-17.7 0-32 14.3-32 32s14.3 32 32 32h32v64H32c-17.7 0-32 14.3-32 32s14.3 32 32 32h576c17.7 0 32-14.3 32-32s-14.3-32-32-32h-32v-64h32c17.7 0 32-14.3 32-32s-14.3-32-32-32H32zm96 64h160v64H128v-64zm224 0h160v64H352v-64z"/>`),
		g.Group(children),
	)
}
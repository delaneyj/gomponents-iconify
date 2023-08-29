package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrancSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M80 32c-17.7 0-32 14.3-32 32v256H32c-17.7 0-32 14.3-32 32s14.3 32 32 32h16v64c0 17.7 14.3 32 32 32s32-14.3 32-32v-64h80c17.7 0 32-14.3 32-32s-14.3-32-32-32h-80v-64h144c17.7 0 32-14.3 32-32s-14.3-32-32-32H112V96h176c17.7 0 32-14.3 32-32s-14.3-32-32-32H80z"/>`),
		g.Group(children),
	)
}
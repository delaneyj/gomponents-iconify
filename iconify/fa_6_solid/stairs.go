package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stairs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384 64c0-17.7 14.3-32 32-32h128c17.7 0 32 14.3 32 32s-14.3 32-32 32h-96v96c0 17.7-14.3 32-32 32h-96v96c0 17.7-14.3 32-32 32h-96v96c0 17.7-14.3 32-32 32H32c-17.7 0-32-14.3-32-32s14.3-32 32-32h96v-96c0-17.7 14.3-32 32-32h96v-96c0-17.7 14.3-32 32-32h96V64z"/>`),
		g.Group(children),
	)
}
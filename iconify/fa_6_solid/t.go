package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func T(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 32C14.3 32 0 46.3 0 64s14.3 32 32 32h128v352c0 17.7 14.3 32 32 32s32-14.3 32-32V96h128c17.7 0 32-14.3 32-32s-14.3-32-32-32H32z"/>`),
		g.Group(children),
	)
}
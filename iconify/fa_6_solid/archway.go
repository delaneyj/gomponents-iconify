package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 32C14.3 32 0 46.3 0 64s14.3 32 32 32h448c17.7 0 32-14.3 32-32s-14.3-32-32-32H32zm0 384c-17.7 0-32 14.3-32 32s14.3 32 32 32h128V352c0-53 43-96 96-96s96 43 96 96v128h128c17.7 0 32-14.3 32-32s-14.3-32-32-32V128H32v288z"/>`),
		g.Group(children),
	)
}
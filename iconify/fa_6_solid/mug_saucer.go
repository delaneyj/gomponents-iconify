package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MugSaucer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M96 64c0-17.7 14.3-32 32-32h384c70.7 0 128 57.3 128 128s-57.3 128-128 128h-32c0 53-43 96-96 96H192c-53 0-96-43-96-96V64zm384 160h32c35.3 0 64-28.7 64-64s-28.7-64-64-64h-32v128zM32 416h512c17.7 0 32 14.3 32 32s-14.3 32-32 32H32c-17.7 0-32-14.3-32-32s14.3-32 32-32z"/>`),
		g.Group(children),
	)
}
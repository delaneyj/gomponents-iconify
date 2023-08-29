package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationCrosshairs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0c17.7 0 32 14.3 32 32v34.7c80.4 13.4 143.9 76.9 157.3 157.3H480c17.7 0 32 14.3 32 32s-14.3 32-32 32h-34.7c-13.4 80.4-76.9 143.9-157.3 157.3V480c0 17.7-14.3 32-32 32s-32-14.3-32-32v-34.7C143.6 431.9 80.1 368.4 66.7 288H32c-17.7 0-32-14.3-32-32s14.3-32 32-32h34.7C80.1 143.6 143.6 80.1 224 66.7V32c0-17.7 14.3-32 32-32zM128 256a128 128 0 1 0 256 0a128 128 0 1 0-256 0zm128-80a80 80 0 1 1 0 160a80 80 0 1 1 0-160z"/>`),
		g.Group(children),
	)
}
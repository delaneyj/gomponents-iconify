package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func J(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M288 32c17.7 0 32 14.3 32 32v256c0 88.4-71.6 160-160 160S0 408.4 0 320v-32c0-17.7 14.3-32 32-32s32 14.3 32 32v32c0 53 43 96 96 96s96-43 96-96V64c0-17.7 14.3-32 32-32z"/>`),
		g.Group(children),
	)
}
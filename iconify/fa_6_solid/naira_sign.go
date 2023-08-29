package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NairaSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M122.6 46.3c-7.8-11.7-22.4-17-35.9-12.9S64 49.9 64 64v192H32c-17.7 0-32 14.3-32 32s14.3 32 32 32h32v128c0 17.7 14.3 32 32 32s32-14.3 32-32V320h100.2l97.2 145.8c7.8 11.7 22.4 17 35.9 12.9s22.7-16.5 22.7-30.6V320h32c17.7 0 32-14.3 32-32s-14.3-32-32-32h-32V64c0-17.7-14.3-32-32-32s-32 14.3-32 32v192h-57.5L122.6 46.3zM305.1 320H320v22.3L305.1 320zm-119.6-64H128v-86.3l57.5 86.3z"/>`),
		g.Group(children),
	)
}
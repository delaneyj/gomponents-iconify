package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RubElHizb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 16.97l-6.4 6.37L185.9 87H87v99l-70.03 70L87 326v99h99l70 70l70-70h99v-99l70-70l-70-70V87h-99l-70-70.03zm0 25.46L300.6 87h-89.2L256 42.43zM105 105h63l-63 63v-63zm88.4 0h125.2l88.4 88.4v125.2L318.6 407H193.4L105 318.6V193.4l88.4-88.4zm150.6 0h63v63l-63-63zm-88 78c-40.2 0-73 32.8-73 73s32.8 73 73 73s73-32.8 73-73s-32.8-73-73-73zm0 18c30.5 0 55 24.5 55 55s-24.5 55-55 55s-55-24.5-55-55s24.5-55 55-55zm169 10.4l44.6 44.6l-44.6 44.6v-89.2zm-338 0v89.2L42.43 256L87 211.4zM407 344v63h-63l63-63zm-302 0l63 63h-63v-63zm106.4 81h89.2L256 469.6L211.4 425z"/>`),
		g.Group(children),
	)
}
package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YenSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M58.6 46.2C48.8 31.5 29 27.6 14.3 37.4S-4.4 67 5.4 81.7L100.2 224H48c-17.7 0-32 14.3-32 32s14.3 32 32 32h80v32H48c-17.7 0-32 14.3-32 32s14.3 32 32 32h80v64c0 17.7 14.3 32 32 32s32-14.3 32-32v-64h80c17.7 0 32-14.3 32-32s-14.3-32-32-32h-80v-32h80c17.7 0 32-14.3 32-32s-14.3-32-32-32h-52.2l94.8-142.3c9.8-14.7 5.8-34.6-8.9-44.4s-34.6-5.8-44.4 8.9L160 198.3L58.6 46.2z"/>`),
		g.Group(children),
	)
}
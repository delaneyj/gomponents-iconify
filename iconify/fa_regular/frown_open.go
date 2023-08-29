package fa_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrownOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M248 8C111 8 0 119 0 256s111 248 248 248s248-111 248-248S385 8 248 8zm0 448c-110.3 0-200-89.7-200-200S137.7 56 248 56s200 89.7 200 200s-89.7 200-200 200zm-48-248c0-17.7-14.3-32-32-32s-32 14.3-32 32s14.3 32 32 32s32-14.3 32-32zm128-32c-17.7 0-32 14.3-32 32s14.3 32 32 32s32-14.3 32-32s-14.3-32-32-32zm-80 112c-35.6 0-88.8 21.3-95.8 61.2c-2 11.8 9 21.5 20.5 18.1c31.2-9.6 59.4-15.3 75.3-15.3s44.1 5.7 75.3 15.3c11.4 3.5 22.5-6.3 20.5-18.1c-7-39.9-60.2-61.2-95.8-61.2z"/>`),
		g.Group(children),
	)
}
package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M409 672q33 0 52 26q266 364 362 774H377Q250 1031 10 723q-12-16-3-33.5T36 672h373zm559 357q-49 199-125 393q-79-310-256-594q40-221 44-449q211 340 337 650zm99-709q235 324 384.5 698.5T1636 1792h-451q-41-665-553-1472h435zm693 576q0 424-101 812q-67-560-359-1083q-25-301-106-584q-4-16 5.5-28.5T1225 0h359q21 0 38.5 13t22.5 33q115 409 115 850z"/>`),
		g.Group(children),
	)
}
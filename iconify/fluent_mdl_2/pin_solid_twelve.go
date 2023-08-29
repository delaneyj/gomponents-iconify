package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinSolidTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1963 512h85v853h-85q-103 0-190-44t-150-126h-317q-37 78-93 141t-127 107t-151 69t-167 24h-85v-512H171L0 939l171-86h512V341h85q86 0 167 24t151 69t127 108t93 141h317q62-81 149-126t191-45z"/>`),
		g.Group(children),
	)
}
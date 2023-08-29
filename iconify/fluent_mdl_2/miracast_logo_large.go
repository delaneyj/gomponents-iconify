package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MiracastLogoLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v1152h-768v-128h640V384H384v256H256V256h1792zM1024 1280h128v256q0 27-10 50t-27 40t-41 28t-50 10H128q-27 0-50-10t-40-27t-28-41t-10-50V896q0-27 10-50t27-40t41-28t50-10h512v128H128v640h896v-256zM768 768q80 0 150 30t122 82t82 122t30 150h-128q0-53-20-99t-55-82t-81-55t-100-20V768zm0 256q27 0 50 10t40 27t28 41t10 50H768v-128zm512 128q0-106-40-199t-110-162t-163-110t-199-41V512q133 0 249 50t204 137t137 203t50 250h-128z"/>`),
		g.Group(children),
	)
}
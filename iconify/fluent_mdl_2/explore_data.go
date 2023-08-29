package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExploreData(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1408v-256H0V0h1536v256h256v512h-128V384H384v896h768v128H256zm0-384V256h1152V128H128v896h128zm256 128V896h128v256H512zm256 0V640h128v512H768zm256 0V768h128v384h-128zm640-256q79 0 149 30t122 82t83 123t30 149q0 80-30 149t-82 122t-123 83t-149 30q-60 0-117-18t-105-53l-437 436q-19 19-45 19t-45-19t-19-45q0-26 19-45l436-437q-35-48-53-105t-18-117q0-79 30-149t82-122t122-83t150-30zm0 640q53 0 99-20t82-55t55-81t20-100q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100q0 53 20 99t55 82t81 55t100 20z"/>`),
		g.Group(children),
	)
}
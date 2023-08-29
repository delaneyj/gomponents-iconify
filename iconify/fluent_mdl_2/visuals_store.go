package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VisualsStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1152 0q79 0 149 30t122 82t83 123t30 149v128h256v640h-128V640h-256v768h-128V640H256v1152q0 27 10 50t27 40t41 28t50 10h640v128H384q-53 0-99-20t-82-55t-55-81t-20-100V512h256V384q0-79 30-149t82-122t122-83T768 0q104 0 193 52q89-52 191-52zm256 384q0-53-20-99t-55-82t-81-55t-100-20q-45 0-85 15q29 36 46 71t25 70t11 71t3 77v80h256V384zM512 512h512V384q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100v128zm1536 896v640h-128v-640h128zm-384-128h128v768h-128v-768zm-256 256h128v512h-128v-512zm-256 256h128v256h-128v-256z"/>`),
		g.Group(children),
	)
}
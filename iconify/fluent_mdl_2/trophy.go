package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 256v447q0 57-19 109t-54 94t-83 71t-104 40q-9 75-41 140t-83 117t-116 85t-140 44v133h384v256h128v128H384v-128h128v-256h384v-133q-75-10-140-43t-115-85t-83-117t-42-141q-56-11-104-40t-82-70t-54-94t-20-110V256h256V128h896v128h256zM640 1664v128h640v-128H640zM384 703q0 30 9 58t26 52t40 42t53 29V384H384v319zm576 577q66 0 124-25t101-68t69-102t26-125V256H640v704q0 66 25 124t68 102t102 69t125 25zm576-896h-128v500q28-10 52-28t40-42t26-52t10-59V384z"/>`),
		g.Group(children),
	)
}
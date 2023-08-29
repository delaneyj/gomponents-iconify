package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PivotChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 0v2048H0V0h2048zm-128 128H128v1792h1792V128zM640 640H256V256h384v384zM512 384H384v128h128V384zm128 1408H256V768h384v1024zM512 896H384v768h128V896zm1280-256H768V256h1024v384zm-128-256H896v128h768V384zM979 1709l-237-237l237-237l90 90l-82 83h293q27 0 50-10t40-27t28-41t10-50V987l-83 82l-90-90l237-237l237 237l-90 90l-83-82v293q0 53-20 99t-55 82t-81 55t-100 20H987l82 83l-90 90z"/>`),
		g.Group(children),
	)
}
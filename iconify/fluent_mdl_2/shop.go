package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 512v1280q0 53-20 99t-55 81t-82 55t-99 21H384q-53 0-99-20t-81-55t-55-81t-21-100V512h256V384q0-79 30-149t83-122t122-82T768 0q104 0 193 52q89-52 191-52q79 0 149 30t122 83t82 122t31 149v128h256zm-384-128q0-52-20-99t-55-81t-82-55t-99-21q-45 0-85 15q29 36 46 71t25 70t11 71t3 77v80h256V384zM512 512h512V384q0-52-20-99t-55-81t-82-55t-99-21q-53 0-99 20t-81 55t-55 82t-21 99v128zm802 1408q-34-60-34-128V640H256v1152q0 27 10 50t27 40t41 28t50 10h930zm350-1280h-256v1152q0 27 10 50t27 40t41 28t50 10q27 0 50-10t40-27t28-41t10-50V640z"/>`),
		g.Group(children),
	)
}
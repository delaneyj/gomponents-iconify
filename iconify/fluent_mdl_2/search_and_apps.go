package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchAndApps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 896q79 0 149 30t122 83t82 122t31 149q0 79-30 149t-83 122t-122 82t-149 31q-60 0-117-18t-105-53q-16 16-53 55t-87 90t-103 105t-103 100t-85 75t-51 30q-26 0-45-19t-19-45q0-14 29-51t75-85t100-102t105-104t90-86t56-54q-35-48-53-105t-18-117q0-79 30-149t83-122t122-82t149-31zm0 640q52 0 99-20t81-55t55-81t21-100q0-52-20-99t-55-81t-82-55t-99-21q-53 0-99 20t-81 55t-55 82t-21 99q0 53 20 99t55 81t81 55t100 21zM256 1408v-256H0V0h1536v256h256v530q-32-8-64-13t-64-5V384H384v896h768q0 32 5 64t13 64H256zm0-384V256h1152V128H128v896h128z"/>`),
		g.Group(children),
	)
}
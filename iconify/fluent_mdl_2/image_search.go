package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 576q0 26-19 45t-45 19q-26 0-45-19t-19-45q0-26 19-45t45-19q26 0 45 19t19 45zm384-320v639q-28-28-60-50t-68-42V384H128v677l448-447l556 555q-12 54-12 111q0 16 2 32t4 32L576 794l-448 449v421h835l-128 128H0V256h2048zm-384 640q79 0 149 30t122 82t83 123t30 149q0 80-30 149t-82 122t-123 83t-149 30q-60 0-116-18t-106-54l-437 437q-19 19-45 19t-45-19t-19-45q0-26 19-45l437-437q-35-49-53-105t-19-117q0-79 30-149t82-122t122-83t150-30zm0 640q53 0 99-20t82-55t55-81t20-100q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100q0 53 20 99t55 82t81 55t100 20z"/>`),
		g.Group(children),
	)
}
package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SIPMove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m493 749l-274 275l274 275l-90 90l-366-365l366-365l90 90zm806 806l90 90l-365 366l-365-366l90-90l275 274l275-274zm712-531l-366 365l-90-90l274-275l-274-275l90-90l366 365zM749 493l-90-90l365-366l365 366l-90 90l-275-274l-275 274zm275 147q79 0 149 30t122 82t83 123t30 149q0 80-30 149t-82 122t-123 83t-149 30q-80 0-149-30t-122-82t-83-122t-30-150q0-79 30-149t82-122t122-83t150-30zm0 640q53 0 99-20t82-55t55-81t20-100q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100q0 53 20 99t55 82t81 55t100 20z"/>`),
		g.Group(children),
	)
}
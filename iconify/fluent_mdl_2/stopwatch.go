package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 1152v128H896V640h128v512h384zm231-652q135 136 208 297t73 355q0 124-32 238t-90 214t-140 181t-181 140t-214 91t-239 32q-124 0-238-32t-214-90t-181-140t-140-181t-91-214t-32-239q0-111 26-216t75-198t118-172t154-141t185-103t210-57V128H640V0h640v128h-256v128q139 0 270 41t245 122l208-208l90 90l-198 199zm-615 1420q159 0 298-60t244-165t165-244t61-299q0-159-60-298t-165-244t-244-165t-299-61q-159 0-298 60T482 609T317 853t-61 299q0 159 60 298t165 244t244 165t299 61z"/>`),
		g.Group(children),
	)
}
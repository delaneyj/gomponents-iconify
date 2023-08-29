package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1297 38q166 45 304 140t237 226t155 289t55 331q0 141-36 272t-103 245t-160 207t-208 160t-245 103t-272 37q-141 0-272-36t-245-103t-207-160t-160-208t-103-244t-37-273q0-140 37-272t105-248t167-212t221-164H256V0h512v512H640V215q-117 56-211 140T267 545T164 773t-36 251q0 123 32 237t90 214t141 182t181 140t214 91t238 32q123 0 237-32t214-90t182-141t140-181t91-214t32-238q0-150-48-289t-136-253t-207-197t-266-124l34-123z"/>`),
		g.Group(children),
	)
}
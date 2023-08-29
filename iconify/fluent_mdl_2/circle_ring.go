package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleRing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 0q141 0 272 36t244 104t207 160t161 207t103 245t37 272q0 141-36 272t-104 244t-160 207t-207 161t-245 103t-272 37q-141 0-272-36t-244-104t-207-160t-161-207t-103-245t-37-272q0-141 36-272t104-244t160-207t207-161T752 37t272-37zm0 1920q124 0 238-32t214-90t181-140t140-181t91-214t32-239q0-124-32-238t-90-214t-140-181t-181-140t-214-91t-239-32q-124 0-238 32t-214 90t-181 140t-140 181t-91 214t-32 239q0 124 32 238t90 214t140 181t181 140t214 91t239 32z"/>`),
		g.Group(children),
	)
}
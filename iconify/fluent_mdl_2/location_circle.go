package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 128q123 0 237 32t214 90t182 141t140 181t91 214t32 238q0 123-32 237t-90 214t-141 182t-181 140t-214 91t-238 32q-123 0-237-32t-214-90t-182-141t-140-181t-91-214t-32-238q0-123 32-237t90-214t141-182t181-140t214-91t238-32zm0 1664q106 0 204-27t183-78t156-120t120-155t77-184t28-204q0-106-27-204t-78-183t-120-156t-155-120t-184-77t-204-28q-106 0-204 27t-183 78t-156 120t-120 155t-77 184t-28 204q0 106 27 204t78 183t120 156t155 120t184 77t204 28z"/>`),
		g.Group(children),
	)
}
package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1504 127q113 0 212 42t173 117t116 173t43 212q0 109-41 208t-118 176l-865 865l-865-865Q83 979 42 880T0 671q0-113 42-212t117-173t173-116t212-43q111 0 208 40t177 119q24 23 47 47t48 48q24-24 47-48t48-47q79-78 176-118t209-41zm294 838q59-59 90-135t31-159q0-87-32-162t-89-132t-131-89t-163-33q-84 0-159 31t-135 90q-66 65-128 131T955 639l384 384l-302 301l-90-90l211-211l-384-384l163-164q-26-23-50-49t-49-50q-60-59-134-90t-160-31q-87 0-162 32t-132 89t-89 132t-33 163q0 82 32 159t90 135l774 774l774-774z"/>`),
		g.Group(children),
	)
}
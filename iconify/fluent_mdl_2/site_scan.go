package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SiteScan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1031q-28-28-60-50t-68-40V640H128v1024h1088l-128 128H0V128h2048zm-128 384V256H128v256h1792zm-509 954q0-66 25-123t68-99t100-67t124-24q67 0 125 25t101 68t68 102t25 125q0 65-24 122t-67 101t-99 68t-123 25q-48 0-94-13t-87-39l-201 201q-23 23-45 46t-46 45q-10 9-19 14t-24 5q-26 0-46-20t-20-46q0-14 5-23t14-20q21-24 44-46t47-45q51-51 100-101t101-101q-25-41-38-86t-14-94zm128 12q0 38 14 71t39 58t58 39t72 14q39 0 74-15t64-43t44-62t17-74q0-36-15-69t-42-60t-60-42t-70-16q-39 0-74 16t-62 45t-43 63t-16 75z"/>`),
		g.Group(children),
	)
}
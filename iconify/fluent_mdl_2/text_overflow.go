package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextOverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m507 650l234 628H626l-56-160H323l-54 160H155l235-628h117zm36 383l-88-250q-3-9-5-20t-4-20q-2 9-3 19t-6 21l-87 250h193zm498-214q48 0 82 18t57 50t32 71t10 83q0 48-11 92t-36 79t-62 55t-92 21q-85 0-134-75v65H785V614h104v294q51-89 152-89zm-44 389q34 0 57-15t38-40t20-53t6-59q0-26-6-51t-20-46t-35-32t-51-12q-29 0-51 11t-38 29t-23 43t-8 52v55q0 24 8 45t22 38t35 25t46 10zm987-56q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45q0-26 19-45t45-19zm-256 0q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45q0-26 19-45t45-19zm-256 0q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45q0-26 19-45t45-19z"/>`),
		g.Group(children),
	)
}
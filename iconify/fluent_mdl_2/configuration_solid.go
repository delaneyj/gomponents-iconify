package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConfigurationSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1755 512h-475V37l475 475zm-795 520q38 0 71 14t59 40t39 59t15 71q0 38-14 71t-40 59t-59 39t-71 15q-38 0-71-14t-59-40t-39-59t-15-71q0-38 14-71t40-59t59-39t71-15zm832-392v1408H128V0h1024v640h640zm-509 632q2-14 3-28t1-28q0-14-1-28t-3-28l185-76l-55-134l-185 77q-33-46-79-79l77-185l-134-55l-76 185q-14-2-28-3t-28-1q-14 0-28 1t-28 3l-76-185l-134 55l77 185q-46 33-79 79l-185-77l-55 134l185 76q-2 14-3 28t-2 28q0 14 1 28t4 28l-185 76l55 134l185-77q33 46 79 79l-77 185l134 55l76-185q14 2 28 3t28 2q14 0 28-1t28-4l76 185l134-55l-77-185q46-33 79-79l185 77l55-134l-185-76z"/>`),
		g.Group(children),
	)
}
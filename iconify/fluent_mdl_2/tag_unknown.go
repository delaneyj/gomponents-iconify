package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagUnknown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1600 448q0 27-10 50t-27 40t-41 28t-50 10q-27 0-50-10t-40-27t-28-41t-10-50q0-27 10-50t27-40t41-28t50-10q27 0 50 10t40 27t28 41t10 50zM181 1024l715 715l384-384v181l-384 384L0 1024L1024 0h896v896l-105 105q-44-9-87-9q-25 0-49 3t-49 10l162-162V128h-715l-896 896zm1483 1024v-128h128v128h-128zm64-896q53 0 99 20t82 55t55 81t20 100q0 46-14 81t-37 65t-52 57t-59 58q-20 20-25 42t-6 48v16q0 8 1 17h-128v-48q0-46 14-81t35-63t47-50t46-45t36-45t14-52q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50h-128q0-53 20-99t55-82t81-55t100-20z"/>`),
		g.Group(children),
	)
}
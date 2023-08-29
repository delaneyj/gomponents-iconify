package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageShared(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1850 1712q46 25 82 61t62 80t40 93t14 102h-128q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100h-128q0-52 14-101t39-94t62-80t83-61q-33-35-51-81t-19-95q0-53 20-99t55-82t81-55t100-20q53 0 99 20t82 55t55 81t20 100q0 49-18 95t-52 81zm-314-176q0 27 10 50t27 40t41 28t50 10q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50zm-384-896V128H256v1792h880q-8 32-12 64t-4 64H128V0h1115l549 549v591q-63-20-128-20V640h-512zm128-128h293l-293-293v293z"/>`),
		g.Group(children),
	)
}
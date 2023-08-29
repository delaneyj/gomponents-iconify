package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1152 640V128H256v1792h896v128H128V0h1115l549 549v337q-31-10-63-16t-65-6V640h-512zm128-128h293l-293-293v293zm768 896v640h-768v-640h128v-128q0-53 20-99t55-82t81-55t100-20q53 0 99 20t82 55t55 81t20 100v128h128zm-512 0h256v-128q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50v128zm384 128h-512v384h512v-384z"/>`),
		g.Group(children),
	)
}
package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1728 512q66 0 124 25t101 69t69 102t26 124v292q-20-50-53-93t-75-76V832q0-40-15-75t-41-61t-61-41t-75-15H320q-40 0-75 15t-61 41t-41 61t-15 75v320h1152v128H0V832q0-66 25-124t68-101t102-69t125-26h1408zm64 256v128h-128V768h128zm-256 128h-128V768h128v128zm128 128q53 0 99 20t82 55t55 81t20 100v128h128v640h-768v-640h130v-128q1-53 20-99t54-82t80-55t100-20zm-128 384h256v-128q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50v128zm384 128h-512v384h512v-384z"/>`),
		g.Group(children),
	)
}
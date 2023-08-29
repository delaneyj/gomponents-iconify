package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1673 119q90 90 160 195t117 220t73 239t25 251q0 128-25 251t-72 239t-118 220t-160 195l-91-91q81-81 144-175t106-198t65-215t22-226q0-115-22-226t-65-214t-106-198t-144-176l91-91zm-9 905q0 180-68 343t-194 291l-91-91q109-109 167-249t58-294q0-154-58-294t-167-249l91-91q126 128 194 291t68 343zm-534-362q73 73 111 166t39 196q0 103-38 196t-112 166l-90-90q54-54 83-124t29-148q0-77-29-147t-83-125l90-90zM677 256h91v1536h-90l-385-384H0V640h293l384-384zm-37 219L347 768H128v512h219l293 293V475z"/>`),
		g.Group(children),
	)
}
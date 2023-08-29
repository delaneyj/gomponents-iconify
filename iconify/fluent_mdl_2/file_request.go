package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileRequest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 512v896h-128V512h-704q-56 0-90 9t-58 24t-41 31t-37 31t-50 23t-76 10H128v896h896v128H128q-27 0-50-10t-40-27t-28-41t-10-50V256q0-27 10-50t27-40t41-28t50-10h736q37 0 69 13t58 36t49 51t39 59q13 23 25 41t28 30t35 19t49 7h704q27 0 50 10t40 27t28 41t10 50zm-1184 0q27 0 45-9t35-22t34-28t39-28q-15-17-31-45t-36-56t-40-48t-46-20H128v256h736zm1184 1280q0 53-20 99t-55 82t-81 55t-100 20v-128q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10h-293l162 163l-90 90l-317-317l317-317l90 90l-162 163h293q53 0 99 20t82 55t55 81t20 100z"/>`),
		g.Group(children),
	)
}
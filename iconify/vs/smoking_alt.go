package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmokingAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1504 640q0-23-5-40t-17-27.5t-22.5-17t-33-9.5t-36.5-4.5t-43-2.5q-47-2-72-4t-72-13t-75.5-29.5t-66-56t-61-89.5T957 214.5T928 32q0-14 9-23t23-9t23 9t9 23q7 86 23 154t35.5 114t47.5 79.5t54.5 52t62 29.5t62.5 14t64 4q10 0 35.5.5t35 1t31.5 2.5t31 5t25.5 9t24 14t18 20t15 27t8 36t3.5 46v96q0 14-9 23t-23 9t-23-9t-9-23v-96zm0-320q-2-160-160-160h-96q-14 0-23-9t-9-23t9-23t23-9h96q101 0 158.5 49.5T1560 292q53 0 85 9.5t61 40.5q56 60 54 170v224q0 14-9 23t-23 9t-23-9t-9-23V512q0-72-31-116t-97-44h-64v-32zM0 832h1408v320H0V832zm1503 0q-12 0-21.5 9t-9.5 22v258q0 13 9.5 22t21.5 9h66q12 0 21.5-9t9.5-22V863q0-13-9.5-22t-21.5-9h-66zm192 0q-12 0-21.5 9t-9.5 22v258q0 13 9.5 22t21.5 9h66q12 0 21.5-9t9.5-22V863q0-13-9.5-22t-21.5-9h-66z"/>`),
		g.Group(children),
	)
}
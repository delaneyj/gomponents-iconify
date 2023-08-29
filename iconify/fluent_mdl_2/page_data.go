package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageData(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1216v618q0 41-19 73t-52 55t-74 40t-83 27t-84 14t-72 5q-32 0-73-4t-83-14t-84-26t-73-40t-51-56t-20-74v-618q0-39 20-69t54-51t75-35t84-22t83-12t69-3q29 0 69 3t82 11t84 22t75 35t53 51t20 70zm-128 618v-470q-60 24-125 34t-130 10q-63 0-130-10t-127-34v470q0 8 6 13t12 11q19 17 49 29t64 19t67 11t58 3q26 0 59-3t66-10t63-19t50-30q5-5 11-10t7-14zm0-621q-12-16-46-28t-73-19t-78-10t-58-4q-19 0-58 3t-81 12t-76 22t-43 31q13 17 47 28t74 18t78 11t59 3q19 0 58-3t80-12t74-20t43-32zm-256-573h-512V128H256v1792h896v128H128V0h1115l549 549v347h-128V640zm-384-421v293h293l-293-293z"/>`),
		g.Group(children),
	)
}
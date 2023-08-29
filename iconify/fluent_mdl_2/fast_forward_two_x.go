package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardTwoX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1668 510l-644 520V513l-640 517V-5l640 512V-5l644 515zM512 762l312-251l-312-250v501zm640-501v501l312-251l-312-250zM562 1958h439v90H457v-44q0-25 2-50t10-49q15-55 50-96t79-77t91-68t86-71t64-85t25-110q0-75-43-118t-119-43q-30 0-59 8t-57 24t-51 34t-44 42v-108q47-47 102-66t121-19q54 0 100 16t81 46t53 74t19 101q0 83-30 141t-75 102t-97 78t-98 71t-75 77t-30 100zm970-550h113l-215 324l211 316h-119q-38-64-77-127t-77-129h-2q-8 12-14 24t-15 25l-129 207h-118l218-314l-208-326h119q38 67 76 133t75 136h2l160-269z"/>`),
		g.Group(children),
	)
}
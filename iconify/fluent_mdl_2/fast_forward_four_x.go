package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardFourX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1668 510l-644 520V513l-640 517V-5l640 512V-5l644 515zM512 762l312-251l-312-250v501zm640-501v501l312-251l-312-250zm-243 891v592h115v93H909v211H808v-211H387v-88q59-66 116-137t111-147t101-153t85-160h109zm-101 592v-439q-66 118-143 225t-163 214h306zm724-336h113l-215 324l211 316h-119q-38-64-77-127t-77-129h-2q-8 12-14 24t-15 25l-129 207h-118l218-314l-208-326h119q38 67 76 133t75 136h2l160-269z"/>`),
		g.Group(children),
	)
}
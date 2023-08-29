package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchCommit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 960q0 83-29 158t-80 135t-121 99t-154 51v517H896v-517q-83-11-153-50t-122-99t-80-135t-29-159q0-83 29-158t80-135t121-99t154-51V0h128v517q83 11 153 50t122 99t80 135t29 159zm-448 320q66 0 124-25t101-68t69-102t26-125q0-66-25-124t-69-101t-102-69t-124-26q-66 0-124 25t-102 69t-69 102t-25 124q0 66 25 124t68 102t102 69t125 25z"/>`),
		g.Group(children),
	)
}
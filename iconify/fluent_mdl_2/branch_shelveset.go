package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchShelveset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 1019q-83-11-153-50t-122-99t-80-135t-29-159q0-93 35-174t96-143t142-96t175-35q93 0 174 35t143 96t96 142t35 175q0 83-29 158t-80 135t-121 99t-154 51v517H896v-517zm64-763q-66 0-124 25t-102 69t-69 102t-25 124q0 66 25 124t68 102t102 69t125 25q66 0 124-25t101-68t69-102t26-125q0-66-25-124t-69-101t-102-69t-124-26zm960 896v768H0v-768h128v640h1664v-640h128z"/>`),
		g.Group(children),
	)
}
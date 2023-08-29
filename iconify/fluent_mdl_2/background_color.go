package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackgroundColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v1792H256v-640l128 128v384h1536V384h-768V256h896zM704 1562L102 960l538-539V192q0-40 15-75t41-61t61-41t75-15q40 0 75 15t61 41t41 61t15 75v576H896V192q0-26-19-45t-45-19q-26 0-45 19t-19 45v283L347 896h842l65-64l-83-83l90-90l173 173l-730 730zm357-538H347l357 358l357-358zm417 786q-67 0-129-23t-111-64t-77-100t-29-129q0-43 12-84t35-77l293-461l286 450q23 38 36 79t13 87q0 67-26 126t-72 102t-105 69t-126 25zm-161-389q-22 33-22 71q0 36 15 64t41 49t57 30t65 11q33 0 64-11t54-33t38-50t15-64q0-41-24-79l-148-233l-155 245z"/>`),
		g.Group(children),
	)
}
package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wpbeginner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M384 800h160V576H384v224zm837 332v-92q-104 36-243 38q-135 1-259.5-46.5T498 909l1 96q88 80 212 128.5t272 47.5q129 0 238-49zM640 800h640V576H640v224zm1152-32q0 187-99 352q89 102 89 229q0 157-129.5 268T1339 1728q-122 0-225-52.5T953 1535q-19 1-57 1t-57-1q-58 88-161 140.5T453 1728q-184 0-313.5-111T10 1349q0-127 89-229Q0 955 0 768q0-209 120-385.5T446.5 103T896 0t449.5 103T1672 382.5T1792 768z"/>`),
		g.Group(children),
	)
}
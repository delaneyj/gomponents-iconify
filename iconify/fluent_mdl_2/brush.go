package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1803 128q51 0 95 19t78 53t52 77t20 96q0 49-18 93t-54 80q-303 303-605 606t-606 606q-6 61-33 114t-69 92t-98 61t-117 23H192q-40 0-75-15t-61-41t-41-61t-15-75v-64h64q26 0 45-19t19-45q0-61 22-116t62-98t92-70t114-33q303-303 606-605t606-606q35-35 79-53t94-19zM576 1435q55 24 98 67t67 98l116-116q-61-104-165-165l-116 116zm-128 485q40 0 75-15t61-41t41-61t15-75q0-40-15-75t-41-61t-61-41t-75-15q-40 0-75 15t-61 41t-41 61t-15 75q0 55-29 102t-80 71q18 19 45 19h256zM1886 456q34-34 34-83q0-24-9-45t-25-38t-37-25t-46-9q-49 0-83 34l-935 936q99 66 165 165l936-935z"/>`),
		g.Group(children),
	)
}
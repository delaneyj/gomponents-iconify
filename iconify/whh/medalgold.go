package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medalgold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1015.717 64l-219 382q-79-87-193-115l153-267q15-26 44-45t56-19h129q27 0 36 19t-6 45zm-183 641q0 87-42.5 160.5T673.717 982t-161 43t-161-43t-116.5-116.5t-42.5-160.5t42.5-161t116.5-117t161-43t161 43t116.5 117t42.5 161zm-224 128h-32V640q0-53-37.5-90.5t-90.5-37.5v1v-1q-14 0-23 9.5t-9 23t9 23t23 9.5v-1q27 1 45.5 19.5t18.5 44.5v193h-32q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5h128q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5zm-599-769q-15-26-6-45t36-19h129q27 0 56 19t44 45l153 267q-114 28-193 115z"/>`),
		g.Group(children),
	)
}
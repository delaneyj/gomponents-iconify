package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medalsilver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1015.717 64l-219 382q-79-87-193-114l153-268q15-26 44-45t56-19h129q27 0 36 19t-6 45zm-183 641q0 87-42.5 160.5T673.717 982t-161 43t-161-43t-116.5-116.5t-42.5-160.5t42.5-161t116.5-117t161-43t161 43t116.5 117t42.5 161zm-320 64q53 0 90.5-37.5t37.5-90.5t-37.5-91t-90.5-38t-90.5 38t-37.5 91q0 13 9 22.5t22.5 9.5t23-9.5t9.5-22.5q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45t-19 45.5t-45 19q-53 0-90.5 37.5t-37.5 90.5v32q0 13 9 22.5t23 9.5h192q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-160q0-27 18.5-45.5t45.5-18.5zm-503-705q-15-26-6-45t36-19h129q27 0 56 19t44 45l153 268q-114 27-193 114z"/>`),
		g.Group(children),
	)
}
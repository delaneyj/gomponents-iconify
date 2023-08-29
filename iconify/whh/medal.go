package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1015.717 64l-219 382q-80-87-193-114l153-268q15-26 44-45t56-19h129q27 0 36 19t-6 45zm-183 641q0 87-42.5 160.5T673.717 982t-161 43t-161-43t-116.5-116.5t-42.5-160.5t42.5-161t116.5-117t161-43t161 43t116.5 117t42.5 161zm-823-641q-15-26-6-45t36-19h129q27 0 56 19t44 45l153 268q-114 27-193 114z"/>`),
		g.Group(children),
	)
}
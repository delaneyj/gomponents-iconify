package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmbedTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m13 11.5l1.5 1.5l5-5l-5-5L13 4.5L16.5 8zm-6-7L5.5 3l-5 5l5 5L7 11.5L3.5 8zm3.958-2.148l1.085.296l-3 11l-1.085-.296l3-11z"/>`),
		g.Group(children),
	)
}
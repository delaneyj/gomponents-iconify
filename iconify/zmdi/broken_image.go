package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrokenImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M384 43v140l-64-64l-85 86l-86-86l-85 86l-64-65V43q0-18 12.5-30.5T43 0h298q18 0 30.5 12.5T384 43zm-64 137l64 64v97q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V201l64 64l85-86l86 86z"/>`),
		g.Group(children),
	)
}
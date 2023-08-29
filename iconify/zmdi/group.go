package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Group(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M149 155v85h107v-85H149zM0 5h85v22h299V5h85v86h-21v298h21v86h-85v-22H85v22H0v-86h21V91H0V5zm85 384v22h299v-22h21V91h-21V69H85v22H64v298h21zm22-277h192v85h64v171H149v-85h-42V112zm192 171H192v42h128v-85h-21v43z"/>`),
		g.Group(children),
	)
}
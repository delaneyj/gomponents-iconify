package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ungroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 3h85v21h150V3h85v85h-21v64h42v-21h86v85h-22v128h22v85h-86v-21H213v21h-85v-85h21v-43H85v22H0v-86h21V88H0V3zm341 213v-21h-42v42h21v86h-85v-22h-43v43h21v21h128v-21h22V216h-22zM235 88V67H85v21H64v149h21v22h64v-43h-21v-85h85v21h43V88h-21zm-22 128h-21v43h43v-22h21v-42h-43v21z"/>`),
		g.Group(children),
	)
}
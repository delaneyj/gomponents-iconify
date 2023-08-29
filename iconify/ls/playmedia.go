package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playmedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 359c0 99 39 189 105 254c65 65 154 104 253 104s188-39 254-104c65-65 105-155 105-254s-40-189-105-254C546 40 457 0 358 0S170 40 105 105C39 170 0 260 0 359zm108 0c0-70 28-132 73-177s107-74 177-74s133 29 178 74c46 45 73 107 73 177s-27 132-73 177c-45 45-107 74-178 74c-70 0-132-29-177-74s-73-107-73-177zm153-140v281c0 5 2 9 7 12c2 1 5 2 7 2s6-1 8-2l244-141c4-2 8-6 8-12s-4-10-8-12L283 205c-3-2-9-2-15 0c-5 3-7 9-7 14z"/>`),
		g.Group(children),
	)
}
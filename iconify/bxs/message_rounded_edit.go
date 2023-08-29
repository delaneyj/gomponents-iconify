package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageRoundedEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 5.589 2 10c0 2.907 1.897 5.516 5 6.934V22l5.34-4.005C17.697 17.854 22 14.32 22 10c0-4.411-4.486-8-10-8zM9.302 13.986H7.503v-1.798l4.976-4.97l1.798 1.799l-4.975 4.969zm5.823-5.816l-1.799-1.798l1.372-1.371l1.799 1.798l-1.372 1.371z"/>`),
		g.Group(children),
	)
}
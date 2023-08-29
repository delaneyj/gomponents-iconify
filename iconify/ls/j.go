package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func J(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M188 82h-82V0h82v82zm-77 694V164h73v610c0 53-30 99-73 124c-20 12-45 19-72 19c-13 0-27-2-39-5v-78c11 7 25 12 39 12c39 0 71-32 72-70z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CabinAFrame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3L4 21h16L12 3M9 19H7.08L9 14.67V19m4 0h-2v-5h2v5m-2.81-7L12 7.92L13.81 12h-3.62M15 14.67L16.92 19H15v-4.33Z"/>`),
		g.Group(children),
	)
}
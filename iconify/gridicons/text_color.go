package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19h18v3H3v-3zm12.82-2h3.424L14 3h-4L4.756 17h3.425l1.066-3.5h5.506L15.82 17zm-1.952-6h-3.73l1.868-5.725L13.868 11z"/>`),
		g.Group(children),
	)
}
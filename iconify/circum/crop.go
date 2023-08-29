package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.624 6.623H3.549a.5.5 0 0 1 0-1h2.075V3.55a.5.5 0 0 1 1 0v2.073h9.191a2.562 2.562 0 0 1 2.561 2.561v9.193h2.075a.5.5 0 0 1 0 1h-2.075v2.073a.5.5 0 0 1-1 0v-2.073H8.185a2.562 2.562 0 0 1-2.561-2.561V6.623Zm11.752 10.754V8.184c0-.862-.699-1.561-1.561-1.561H6.624v9.193c0 .862.699 1.561 1.561 1.561h9.191Z"/>`),
		g.Group(children),
	)
}
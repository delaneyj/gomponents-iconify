package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.245 17.666L10 13.467l3.755 4.2c.611.684 1.745.251 1.745-.667V3a1 1 0 0 0-1-1h-9a1 1 0 0 0-1 1v14c0 .918 1.134 1.35 1.745.666Zm.255-3.285V4h7v10.381l-2.755-3.08a1 1 0 0 0-1.49 0L6.5 14.38Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
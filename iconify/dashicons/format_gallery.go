package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatGallery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 4h1.96c.57 0 1.04.47 1.04 1.04v12.92c0 .57-.47 1.04-1.04 1.04H5.04C4.47 19 4 18.53 4 17.96V16H2.04C1.47 16 1 15.53 1 14.96V2.04C1 1.47 1.47 1 2.04 1h12.92c.57 0 1.04.47 1.04 1.04V4zM3 14h11V3H3v11zm5-8.5C8 4.67 7.33 4 6.5 4S5 4.67 5 5.5S5.67 7 6.5 7S8 6.33 8 5.5zm2 4.5s1-5 3-5v8H4V7c2 0 2 3 2 3s.33-2 2-2s2 2 2 2zm7 7V6h-1v8.96c0 .57-.47 1.04-1.04 1.04H6v1h11z"/>`),
		g.Group(children),
	)
}
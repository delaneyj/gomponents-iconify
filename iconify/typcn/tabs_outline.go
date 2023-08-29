package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 4H8a2 2 0 0 0-2 2v2H5a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1h2a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zM5 19v-9h8.5c.275 0 .5.225.5.5V19H5zm13-3h-3v-5.5c0-.827-.673-1.5-1.5-1.5H8V6h10v10z"/>`),
		g.Group(children),
	)
}
package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 13v9m0 0l3.5-3.5M12 22l-3.5-3.5m11.5-.893c1.494-.585 3-1.918 3-4.607c0-4-3.333-5-5-5c0-2 0-6-6-6S6 6 6 8c-1.667 0-5 1-5 5c0 2.689 1.506 4.022 3 4.607"/>`),
		g.Group(children),
	)
}
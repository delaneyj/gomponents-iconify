package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3.5V3m-6.5 7H3m11.596-4.596l.354-.354m-9.546.354L5.05 5.05M9 13.036a3.484 3.484 0 0 1 1.975.99M12.645 7a4 4 0 0 0-6.59 3.666M5 16.5A3.5 3.5 0 0 0 8.5 20h9.75a2.75 2.75 0 0 0 .734-5.4a5 5 0 0 0-9.614-1.491A3.504 3.504 0 0 0 5 16.5z"/>`),
		g.Group(children),
	)
}
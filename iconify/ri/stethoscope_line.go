package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StethoscopeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 3v2H6v4a4 4 0 0 0 8 0V5h-2V3h3a1 1 0 0 1 1 1v5a6.002 6.002 0 0 1-5 5.917V16.5a3.5 3.5 0 0 0 6.775 1.238a3 3 0 1 1 2.05.148A5.502 5.502 0 0 1 8.999 16.5v-1.583A6.002 6.002 0 0 1 4 9V4a1 1 0 0 1 1-1h3Zm11 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}
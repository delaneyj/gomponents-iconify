package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KickScooter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8.014 5H12a1 1 0 1 1 0 2H8.905l-.998 9.983l5.666 6.918A3 3 0 0 0 15.894 25h6.232A4.002 4.002 0 0 1 30 26a4 4 0 0 1-7.874 1h-6.232a5 5 0 0 1-3.868-1.832l-4.4-5.373l-.245 2.45A4.002 4.002 0 0 1 6 30a4 4 0 0 1-.61-7.954L7.004 5.911a.994.994 0 0 1 .65-.85a.994.994 0 0 1 .36-.061ZM5.177 24.177a2 2 0 1 0 1.99.199l-.172 1.723a1 1 0 0 1-1.99-.199l.172-1.723ZM24.267 27a2 2 0 1 0 0-2H26a1 1 0 1 1 0 2h-1.733Z"/>`),
		g.Group(children),
	)
}
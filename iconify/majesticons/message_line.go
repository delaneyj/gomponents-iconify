package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-4.82a1 1 0 0 0-.61.207l-4.35 3.347C7.903 22.566 6 21.63 6 19.97V19a1 1 0 0 0-1-1a3 3 0 0 1-3-3V6zm3-1a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1a3 3 0 0 1 3 3v.97l4.351-3.348a3 3 0 0 1 1.83-.622H19a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H5z"/></g>`),
		g.Group(children),
	)
}